package api

import (
	"Decoder2/logging"
	"bytes"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
)

const smToken = "A12CC37A-9BDD-4466-9610-A822FB935111"
const smUrl = "https://support.wagner-auto.ru/api/v3/requests"

func SmApi(smJson []byte, envVar map[string]string, logger *logrus.Logger) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("input_data", string(smJson))
	if err != nil {
		logger.Panicf("\nUnable to composr input_data JSON, error%v\n", err)
		return
	}
	err = writer.Close()
	if err != nil {
		logger.Panicf("\nUnable to close writer, error:%v\n", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", envVar["APIREQUESTURL"], payload)

	if err != nil {
		logger.Panicf("\nUnable combine POST request, errror:%v\n", err)
		return
	}

	req.Header.Add("authtoken", envVar["APITOKEN"])
	req.Header.Set("Content-Type", writer.FormDataContentType())
	//---
	res, err := client.Do(req)
	if err != nil {
		logger.Panicf("\nUnable to perform POST:%v,%v\n", res, err)
		return
	}
	text, err := logging.HandleApiError(res, logger)
	if err != nil {
		logger.Panicf("\nAPi response code:%v,%v\n", text, err)
		return
	}

	err = res.Body.Close()
	if err != nil {
		logger.Errorf("\nUnable to close body, error=%v\n", err)

	}

}
