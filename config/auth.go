package config

import (
	"time"

	"github.com/sderohan/jwt-authentication-gRPC-go/utils"
)

type AuthConfig struct {
	Username        string
	Password        string
	RefreshDuration time.Duration
	SecretKey       string
}

var authConfig *AuthConfig

func initAuthConfig() {
	authConfig = &AuthConfig{
		Username:        utils.GetString("USERNAME"),
		Password:        utils.GetString("PASSWORD"),
		RefreshDuration: utils.GetDuration("REFRESH_DURATION"),
		SecretKey:       utils.GetString("SECRET_KEY"),
	}
}

func GetAuthConfig() *AuthConfig {
	return authConfig
}
