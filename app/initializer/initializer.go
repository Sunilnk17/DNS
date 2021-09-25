package initializer

import (
	"drone-navigation-service/app/config"
	"drone-navigation-service/app/config/locales/local_config"
	"log"
)

func Initialize() {

	initializeConfig()

}

func initializeConfig() {

	err := config.Initialize()
	local_config.InitializeLocalizer()
	config.InitializeDecoderAndValidator()

	if err != nil {
		log.Fatal(nil, "error initializing Config :", err)
		return
	}

}
