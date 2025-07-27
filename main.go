package main

import (
	"github.com/Lzrb0x/study-go-api.git/config"
	"github.com/Lzrb0x/study-go-api.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization failed: %v", err)
		return
	}

	router.Initialize()
}
