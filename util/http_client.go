package util

import (
	"bytes"
	"io/ioutil"
	"mf_backup_onetime/dto"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CallAPI(data *dto.CallAPIDto) (err error) {
	log.Info("CallAPI() - starting...")
	if err = data.Validate(); err != nil {
		return err
	}

	client := &http.Client{}
	var request *http.Request
	if data.BodyRequest != "" {
		request, err = http.NewRequest(data.Method, data.Url, bytes.NewBuffer([]byte(data.BodyRequest)))
		if err != nil {
			return err
		}
	} else {
		request, err = http.NewRequest(data.Method, data.Url, nil)
		if err != nil {
			return err
		}
	}

	request.Header.Set("Content-Type", data.ContentType)

	if data.Headers != nil && len(data.Headers) > 0 {
		for key, header := range data.Headers {
			request.Header.Set(key, header.(string))
		}
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Errorln("error on close response body: ", err)
		}
	}()

	bodyResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	data.BodyResponse = string(bodyResponse)
	data.HttpCode = response.StatusCode

	log.Info("CallAPI() - finished.")
	return nil
}
