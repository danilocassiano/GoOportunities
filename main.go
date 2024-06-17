package main

import (
	"fmt"

	"github.com/danilocassiano/GoOportunities/config"
	"github.com/danilocassiano/GoOportunities/router"
)

func main() {
	logger := config.GetLogger("main")

	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		fmt.Println(err)
		return
	}

	// Initialize router
	router.Inicialize()
}
