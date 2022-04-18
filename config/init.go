package config

import "github.com/sderohan/jwt-authentication-gRPC-go/utils"

const (
	CONFIG_FILE_NAME = "app"
	CONFIG_FILE_TYPE = "env"
	CONFIG_FILE_PATH = "./.."
)

func setupConfig() {
	utils.SetConfigFileName(CONFIG_FILE_NAME)
	utils.SetConfigFileType(CONFIG_FILE_TYPE)
	utils.SetConfigFileSearchPath(CONFIG_FILE_PATH)
}

func InitConfig() {
	setupConfig()
	initServerConfig()
	initAuthConfig()
}
