package main

import (
	. "Decoder2/controller/api"
	"Decoder2/env"
	"Decoder2/utils"
	"fmt"
	"time"

	"Decoder2/logging"
	"Decoder2/usescases"
	"bufio"
	"encoding/hex"
	"os"
	"strings"
)

const envPath = "conf.env"

func NewDecoderStruct() *DecoderStruct {
	return &DecoderStruct{}
}

func main() {
	Logger := logging.Logger()
	smInfo := NewDecoderStruct()
	smReader := bufio.NewReader(os.Stdin)
	envVar := env.ParseEnv(Logger, envPath)

	for {
		fmt.Print("\n<---- Необходимо приложить карту (Расскладка клавиатуры должна быть английская),  потом нажать Enter ---->\n")
		smReadHex, _ := smReader.ReadString('\n')
		smReadHex = strings.TrimSpace(strings.TrimSuffix(smReadHex, "\n"))
		_, err := hex.DecodeString(smReadHex)

		if len(smReadHex) != 10 {
			Logger.Warnf("\nНеверное количество символов, должно быть 10, а введено: %v\n\n %s\n", len(smReadHex), smReadHex)
		} else if err != nil {
			Logger.Warnf("\n Неверный формат ввода, доступные симоволы 0-9 ABCDEF !!!! \n")
		} else {

			smInfo = &DecoderStruct{
				Bolid: usescases.Bolid(smReadHex, Logger),
				Stork: usescases.Stork(smReadHex, Logger),
				KM:    smReadHex,
				User:  "",
			}

			fmt.Print("\n<---- Введите ФИО, и снова Enter ---->\n")
			smInfo.User, _ = smReader.ReadString('\n')

			utils.ClearScreen(Logger)
			marshaledJSON := SmJson(*smInfo, Logger)
			SmApi(marshaledJSON, envVar, Logger)

			fmt.Printf("\nДанные о карте отправлены в SD\n\nKM: %+v\nStork: %+v\nBolid: %+v\nUser: %+v\nПриложение можно закрывать\n", smInfo.KM, smInfo.Stork, smInfo.Bolid, smInfo.User)

			time.Sleep(1 * time.Minute)
			os.Exit(0)
		}
	}
}
