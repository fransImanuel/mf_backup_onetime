package pkg

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"mf_backup_onetime/schemas"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"

	"strconv"
	"time"
)

type Minio struct {
	Config *schemas.MinioConfig
	Client *minio.Client
}

func InitMinio(config *schemas.MinioConfig) *Minio {
	return &Minio{
		Config: config,
	}
}

func GetMinioConfig(config schemas.SchemaEnvironment) *schemas.MinioConfig {
	log.Info("GetMinioConfig() - starting...")
	minioHost := config.Minio_Host
	minioLocation := config.Minio_Location
	minioAccessKey := config.Minio_AccessKey
	minioSecretKey := config.Minio_SecretKey
	minioSSL := config.Minio_SSL
	minioReplaceDomain := config.Minio_Domain

	bminioSSL, _ := strconv.ParseBool(minioSSL)
	configMinio := &schemas.MinioConfig{
		Host:          minioHost,
		Location:      minioLocation,
		AccessKey:     minioAccessKey,
		SecretKey:     minioSecretKey,
		SSL:           bminioSSL,
		ReplaceDomain: minioReplaceDomain,
	}

	log.Info("GetMinioConfig() - finished.")
	return configMinio
}

func (m *Minio) New() error {
	log.Info("Minio - New() - starting...")
	minioClient, err := minio.New(m.Config.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(m.Config.AccessKey, m.Config.SecretKey, ""),
		Secure: m.Config.SSL,
	})
	if err != nil {
		return err
	}

	m.Client = minioClient

	log.Info("Minio - New() - finished.")
	return nil
}

func (m *Minio) CheckConnection() error {
	if m.Client == nil {
		err := m.New()
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Minio) MakeBucket(name string) error {

	if err := m.CheckConnection(); err != nil {
		return err
	}

	isExist, err := m.Client.BucketExists(context.Background(), name)
	if err != nil {
		return err
	}

	if !isExist {
		err = m.Client.MakeBucket(context.Background(), name, minio.MakeBucketOptions{Region: m.Config.Location})
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Minio) UploadBase64(data *schemas.UploadBase64Dto, bucketName, path string) (err error) {

	if err = data.Validate("upload"); err != nil {
		return err
	}

	decode, err := base64.StdEncoding.DecodeString(data.Base64)
	if err != nil {
		return err
	}

	//fullName := fmt.Sprintf("%v.%v", data.Filename, data.Extension)
	fullName := fmt.Sprintf("%v", data.Filename)
	file, err := os.Create(fullName)
	if err != nil {
		return err
	}

	defer func(fullName string) {
		err := file.Close()
		log.Errorln(err)

		err = os.Remove(fullName)
		if err != nil {
			log.Info("defer UploadBase64, remove file:", fullName, " with error:", err.Error())
		}
	}(fullName)

	if _, err := file.Write(decode); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}

	if err = m.MakeBucket(bucketName); err != nil {
		return err
	}

	uploadInfo, err := m.Client.FPutObject(context.Background(), bucketName, path, fullName, minio.PutObjectOptions{ContentType: data.ContentType})
	if err != nil {
		return err
	}
	data.ResultPath = fullName
	log.Info(fmt.Sprintf("Successfully uploaded %s of size %d\n", fullName, uploadInfo.Size))

	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	err = os.Remove(fullName)
	// 	if err != nil {
	// 		log.Info("error on delete file", fullName, " with error:", err.Error())
	// 	}
	// }()

	fmt.Println("location: ", uploadInfo.Location)

	return nil
}

func (m *Minio) UploadFile(file *os.File, bucketName string) (err error) {
	log.Info("Minio - UploadBase64() - starting...")

	if err = m.MakeBucket(bucketName); err != nil {
		return err
	}

	uploadInfo, err := m.Client.FPutObject(context.Background(), bucketName, file.Name(), file.Name(), minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", file.Name(), uploadInfo.Size)
	fmt.Printf("%+v\n", uploadInfo)
	//data.ResultPath = fullName
	//log.Info(fmt.Sprintf("Successfully uploaded %s of size %d\n", fullName, uploadInfo.Size))
	//
	//err = os.Remove(fullName)
	//if err != nil {
	//	log.Info("error on delete file", fullName, " with error:", err.Error())
	//}

	log.Info("Minio - UploadBase64() - finished.")
	return nil
}

func (m *Minio) DeleteFile(data *schemas.UploadBase64Dto, bucketName string) (err error) {
	log.Info("Minio - DeleteFile() - starting...")
	if err = data.Validate("delete"); err != nil {
		return err
	}

	fullName := fmt.Sprintf("%v.%v", data.Filename, data.Extension)
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err = m.Client.RemoveObject(context.Background(), bucketName, fullName, opts)
	if err != nil {
		return err
	}

	log.Info("Minio - DeleteFile() - finished.")
	return nil
}

func (m *Minio) GetFullUrl(objectName string, bucketName string) (url string, err error) {
	log.Info("Minio - GetFullUrl() - starting...")
	if objectName == "" || bucketName == "" {
		return "", errors.New("request invalid")
	}

	presignedObject, err := m.Client.PresignedGetObject(context.Background(), bucketName, objectName, time.Second*24*60*60, nil)
	if err != nil {
		return "", err
	}

	url = presignedObject.String()

	log.Info("Minio - GetFullUrl() - finished.")
	return url, nil
}
