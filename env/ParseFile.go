package env

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func ParseEnv(logger *logrus.Logger) (envMap map[string]string) {
	envMap, err := godotenv.Read("./env/test.env")
	if err != nil {
		logger.Error("Unable to open .env file")
	}

	return
}
