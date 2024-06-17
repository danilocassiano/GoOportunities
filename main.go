package main

import (
	"fmt"

	"github.com/danilocassiano/GoOportunities/config"
	"github.com/danilocassiano/GoOportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")

	//inicialize config
	err := config.Init()
	if err != nil {	
		logger.Errorf("config initilization error: %v", err)
		fmt.Println(err)
		return 
	}


	//inicialize router
	router.Inicialize()
}