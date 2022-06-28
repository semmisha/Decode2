package logging

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ApiReply struct {
	ResponseStatus struct {
		StatusCode int `json:"status_code"`
		Messages   []struct {
			StatusCode int    `json:"status_code"`
			Type       string `json:"type"`
			Message    string `json:"message"`
		} `json:"messages"`
		Status string `json:"status"`
	} `json:"response_status"`
}

func NewApiReply() *ApiReply {
	return &ApiReply{}
}

var ApiCodeBase = map[int]string{
	200:  "Success",
	2000: "Success",
	4001: "Id or Name given in Input does not exist or not in use or user cannot set the value",
	4002: "Forbidden / User not allowed to perform the operation.",
	4004: "Internal Error (Exact error cannot be sent to user, like some Exception)",
	4005: "Reference Exists. (Cannot delete an entity, because it is being used in another module)",
	4007: "Invalid URL or Resource not found.",
	4008: "Not Unique",
	4009: "Trying to edit non-editable field",
	4012: "Value for mandatory-field is not provided",
	4014: "Trying to edit read-only field",
	4015: "API Rate Limit reached",
	4016: "Time mismatch(Can be used for any time mismatch like Start time/End time mismatch, Created time/Responded time mismatch, etc.,)",
	4021: "Data type mismatch",
	4022: "Invalid API Key",
	7001: "Not allowed as per current license}",
}

func HandleApiError(res *http.Response, logger *logrus.Logger) (string, error) {
	data := NewApiReply()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		logger.Errorf("\nUnable to read body, error:%v\n", err)

	}
	err = json.Unmarshal(read, data)
	if err != nil {
		logger.Errorf("\nUnable to unmarshall, error:%v\n", err)

	}

	for i, _ := range ApiCodeBase {
		if i == data.ResponseStatus.StatusCode {
			if data.ResponseStatus.StatusCode != 2000 {
				return ApiCodeBase[i], fmt.Errorf("%v", data.ResponseStatus.StatusCode)
			}
		}

	}
	return "Success", nil
}
