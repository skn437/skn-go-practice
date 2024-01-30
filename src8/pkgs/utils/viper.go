package utils

import (
	"log"

	"github.com/spf13/viper"
)

func ViperEnvVariable(key string) string {
	viper.SetConfigName("go.config")
	viper.AddConfigPath(".")

	var err error = viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file: %s", err.Error())
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
