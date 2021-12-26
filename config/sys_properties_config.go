package config

import (
	"github.com/flyleft/gprofile"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var AppConfig = ApplicationConfig{}

func InitConfig() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]

	config, err := gprofile.Profile(&ApplicationConfig{}, path + "/application.yaml", true)
	if err != nil {
		log.Fatalf("Profile execute error: %s", err.Error())
	}
	AppConfig = *config.(*ApplicationConfig)

}
