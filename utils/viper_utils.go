package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func GetString(key string) string {
	value := viper.GetString(key)
	if strings.TrimSpace(value) == "" {
		panic(fmt.Sprintf("Key: %s config is not set", key))
	}
	return value
}

func GetDuration(key string) time.Duration {
	value := viper.GetDuration(key)
	if value == 0 {
		panic(fmt.Sprintf("Key: %s config is not set", key))
	}
	return value
}

func SetConfigFileName(filename string) {
	viper.SetConfigFile(filename)
}

func SetConfigFileType(filetype string) {
	viper.SetConfigType(filetype)
}

func SetConfigFileSearchPath(path string) {
	viper.AddConfigPath(path)
}
