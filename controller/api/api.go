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

func SmApi(smJson []byte, logger *logrus.Logger) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("input_data", string(smJson))
	if err != nil {
		logger.Errorf("\nUnable to composr input_data JSON\n", err)
		return
	}
	err = writer.Close()
	if err != nil {
		logger.Errorf("\nUnable to close writer\n", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", smUrl, payload)

	if err != nil {
		logger.Errorf("\nUnable to POST\n", err)
		return
	}
	req.Header.Add("authtoken", smToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	//---
	res, err := client.Do(req)
	text, err := logging.HandleApiError(res, logger)
	if err != nil {
		logger.Errorf("\nAPi response code:%v,%v\n", text, err)
		return
	}

	err = res.Body.Close()
	if err != nil {
		logger.Errorf("\nUnable to close body, error=%v\n", err)

	}

}
