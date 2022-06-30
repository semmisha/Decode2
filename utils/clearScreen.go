package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ClearScreen(logger *logrus.Logger) {

	if strings.Contains(runtime.GOOS, "windows") {

		win := exec.Command("cmd", "/c", "cls")
		win.Stdout = os.Stdout
		if err := win.Run(); err != nil {
			logger.Errorf("\nUnable to clear screen wWndows, error %v\n", err)

		}

	} else if strings.Contains(runtime.GOOS, "linux") || strings.Contains(runtime.GOOS, "darwin") {

		lin := exec.Command("clear")
		lin.Stdout = os.Stdout
		if err := lin.Run(); err != nil {
			logger.Errorf("\nUnable to clear screen Linux, error %v\n", err)

		}
	}

}
