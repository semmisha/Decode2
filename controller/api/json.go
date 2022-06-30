package api

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

type DecoderStruct struct {
	KM    string
	Bolid string
	Stork string
	User  string
}

type SdJson struct {
	Request struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
		Requester   struct {
			Name string `json:"name"`
		} `json:"requester"`
		Status struct {
			Name string `json:"name"`
		} `json:"status"`
		Urgency struct {
			Name string `json:"name"`
		} `json:"urgency"`
		Group struct {
			Name string `json:"name"`
		} `json:"group"`
		Site struct {
			Name string `json:"name"`
		} `json:"site"`
		Priority struct {
			Name string `json:"name"`
		} `json:"priority"`
	} `json:"request"`
}

// Fill the struct/ return []byte
func SmJson(sdInput DecoderStruct, logger *logrus.Logger) []byte {
	var sdData = &SdJson{}
	sdData.Request.Subject = fmt.Sprintf("Новая карта доступа")
	sdData.Request.Description = fmt.Sprintf("%+v", sdInput)
	sdData.Request.Requester.Name = "DCDR"
	sdData.Request.Status.Name = "Открыто"
	sdData.Request.Urgency.Name = "Низкая"
	sdData.Request.Group.Name = "Infrastructure"
	sdData.Request.Site.Name = "ДЦ Пулково"
	sdData.Request.Priority.Name = "Средний"
	var ur, err = json.MarshalIndent(sdData, "", " ")
	if err != nil {
		logger.Panicf("\nUnable to marshall JSON, error:%v\n", err)
	}

	return ur
}
