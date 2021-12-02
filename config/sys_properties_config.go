package config

import (
	"github.com/flyleft/gprofile"
	"log"
)

var AppConfig = ApplicationConfig{}

func InitConfig() {
	config, err := gprofile.Profile(&ApplicationConfig{}, "./application.yaml", true)
	if err != nil {
		log.Fatalf("Profile execute error: %s", err.Error())
	}
	AppConfig = *config.(*ApplicationConfig)

}
