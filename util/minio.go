package util

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"mf_backup_onetime/dto"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
)

type Minio struct {
	Config *dto.MinioConfig
	Client *minio.Client
}

func InitMinio(config *dto.MinioConfig) *Minio {
	return &Minio{
		Config: config,
	}
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
	log.Info("Minio - MakeBucket() - starting...")

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

	log.Info("Minio - MakeBucket() - finished.")
	return nil
}

func (m *Minio) UploadBase64(data *dto.UploadBase64Dto, bucketName string) (err error) {
	log.Info("Minio - UploadBase64() - starting...")
	if err = data.Validate("upload"); err != nil {
		return err
	}

	decode, err := base64.StdEncoding.DecodeString(data.Base64)
	if err != nil {
		return err
	}

	fullName := fmt.Sprintf("%v.%v", data.Filename, data.Extension)
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

	uploadInfo, err := m.Client.FPutObject(context.Background(), bucketName, fullName, fullName, minio.PutObjectOptions{ContentType: data.ContentType})
	if err != nil {
		return err
	}

	data.ResultPath = fullName
	log.Info(fmt.Sprintf("Successfully uploaded %s of size %d\n", fullName, uploadInfo.Size))

	err = os.Remove(fullName)
	if err != nil {
		log.Info("error on delete file", fullName, " with error:", err.Error())
	}

	log.Info("Minio - UploadBase64() - finished.")
	return nil
}

func (m *Minio) DeleteFile(data *dto.UploadBase64Dto, bucketName string) (err error) {
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
