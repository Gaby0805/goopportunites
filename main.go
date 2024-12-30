package main

import (
	"github.com/Gaby0805/goopportunites/config"
	"github.com/Gaby0805/goopportunites/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Inicialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config inicialization error: %v", err)
		return
	}
	//Inicialize Router
	router.Inicialize()
}

//TODO: TERMINAR VIDEO H 1:19
