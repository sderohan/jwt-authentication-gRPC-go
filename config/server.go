package config

import "github.com/sderohan/jwt-authentication-gRPC-go/utils"

type ServerConfig struct {
	Address string
	Port    string
}

var serverConfig *ServerConfig

func initServerConfig() {
	serverConfig = &ServerConfig{
		Address: utils.GetString("ADDRESS"),
		Port:    utils.GetString("PORT"),
	}
}

func GetServerConfig() *ServerConfig {
	return serverConfig
}
