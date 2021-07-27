package config

import (
	"fmt"
	"github.com/flyleft/gprofile"
)

var AppConfig = ApplicationConfig{}

func init() {
	config, err := gprofile.Profile(&ApplicationConfig{}, "./application.yaml", true)
	if err != nil {
		fmt.Errorf("Profile execute error", err)
	}
	AppConfig = *config.(*ApplicationConfig)

}
