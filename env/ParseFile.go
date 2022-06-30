package env

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func ParseEnv(logger *logrus.Logger, filepath string) (envMap map[string]string) {
	_, err := os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(filepath)
		err2 := file.Chmod(777)
		if err2 != nil {
			logger.Errorf("\nCant set Chmod 777 on file:%v , error:%v\n", filepath, err2)
		}
		if err != nil {
			logger.Errorf("\nUnable to create file, error:%v\n", err)
		}
		envText := EnvText()
		_, err = file.Write(envText)
		if err != nil {
			logger.Errorf("\nUnable to write to env file, error:%v\n", err)
		}
		err = file.Close()
		if err != nil {

			logger.Errorf("\nUnable to close .env file, error: %v\n", err)

		}

	} else if err != nil {
		logger.Errorf("Error in os.Stat, filepath:%v , error %v", filepath, err)

	}

	envMap, err = godotenv.Read(filepath)
	if err != nil {
		logger.Panicf("Unable to open .env file")
	}

	return
}
