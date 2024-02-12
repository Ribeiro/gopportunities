package main

import (
	"fmt"

	"github.com/Ribeiro/gopportunities/cmd/api/router"
	"github.com/Ribeiro/gopportunities/config"
)

var (
	logger *config.Logger
)

func main() {
	//Initialize Logger
	logger = config.GetLogger()

	//Initialize Configs
	err := config.Initialize()
	if err != nil {
		fmt.Println(err)
		logger.Errorf("Config initialization failed : %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
