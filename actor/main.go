package main

import (
	"github.com/jpcairesf/cinema/actor/config"
	"github.com/jpcairesf/cinema/actor/router"
)

var (
	logger *config.Logger
)

func main() {
	err := config.Initialize()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	router.Initialize()
	logger = config.GetLogger()
}
