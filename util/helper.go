package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"mf_backup_onetime/constant"
	"mf_backup_onetime/schemas"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
)

func SplitOrderQuery(order string) (string, string) {
	orderA := strings.Split(order, "|")
	if len(orderA) > 1 {
		if orderA[0] != "" && orderA[1] != "" {
			return orderA[0], orderA[1]
		}
	}

	return orderA[0], "ASC"
}
func SplitBasicAuthBase64(token string) (string, string) {

	var decodedByte, _ = base64.StdEncoding.DecodeString(token)
	var decodedString = string(decodedByte)

	orderA := strings.Split(decodedString, ":")
	if len(orderA) > 1 {
		if orderA[0] != "" && orderA[1] != "" {
			return orderA[0], orderA[1]
		}
	}

	return orderA[0], ""
}

func ExcludeCreateSaveDB(columns ...string) []string {
	columnOmit := []string{"created_at", "create_user"}
	if len(columns) > 0 {
		columnOmit = append(columnOmit, columns...)
	}
	return columnOmit
}

func ExcludeCreateDeleteSaveDB(columns ...string) []string {
	columnOmit := []string{"created_at", "create_user", "deleted_at", "deleted_user"}
	if len(columns) > 0 {
		columnOmit = append(columnOmit, columns...)
	}
	return columnOmit
}

func ExcludeDeleteSaveDB(columns ...string) []string {
	columnOmit := []string{"deleted_at", "deleted_user"}
	if len(columns) > 0 {
		columnOmit = append(columnOmit, columns...)
	}
	return columnOmit
}

func HashPassword(password string) (string, error) {
	sha := ""

	h := hmac.New(sha256.New, []byte(constant.KEY_PASSWORD))
	h.Write([]byte(password))
	sha = hex.EncodeToString(h.Sum(nil))

	return sha, nil
}

func HasDuplicates(array []int64) (int64, bool) {
	uniqArray := make(map[int64]bool)
	for _, item := range array {
		if uniqArray[item] {
			return item, true
		}
		uniqArray[item] = true
	}
	return 0, false
}

func ResizeImage(data *multipart.FileHeader, width uint, height uint) (string, error) {
	path := ""

	file, err := data.Open()
	if err != nil {
		log.Println(err)
		return path, err
	}
	img, str, err := image.Decode(file)
	fmt.Println("str ", str)
	if err != nil {
		log.Println(err)
		return path, err
	}
	file.Close()
	m := resize.Resize(width, height, img, resize.Lanczos3)
	if err := os.Mkdir("imageTemp", os.ModePerm); err != nil {
		log.Println(err)
	}

	pathImageThum := "imageTemp/" + data.Filename
	out, err := os.Create(pathImageThum)
	if err != nil {
		log.Println(err)
		return path, err
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
	path = pathImageThum

	return path, err
}

func DeleteFile(data string) error {
	err := os.Remove(data)
	if err != nil {
		log.Println("err_remove_file", err)
	}
	return err
}
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func RemoveDuplicateInt(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	list := []int64{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func APIResponse(ctx *gin.Context, Message string, StatusCode int, Count *int64, Data interface{}) {

	jsonResponse := schemas.SchemaResponses{
		Code: StatusCode,
		//Method:  Method,
		Count:   Count,
		Message: Message,
		Data:    Data,
		//Items:   Items,
	}

	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}
func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}
