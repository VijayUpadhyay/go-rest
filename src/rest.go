package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	config "learn/go-rest/src/config"
	"learn/go-rest/src/router"
	"os"
)

func main() {
	setAppHome()
	configs, err := config.LoadConfig()
	if err != nil {
		logrus.Errorf("Application Load fails while loadng config: %+v ", configs)
	}
	router.Start()
}

func setAppHome() {
	if appHome, err := os.Getwd(); err != nil {
		panic("Failed to get app home path")
	} else {
		err := os.Setenv("APP_HOME", appHome)
		if err != nil {
			panic("Failed to set app home path")
		} else {
			fmt.Println("appHome is set with value: ", appHome)
		}
	}
}
