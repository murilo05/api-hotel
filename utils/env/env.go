package env

import (
	"log"

	"github.com/spf13/viper"
)

func ReadEnvFile() {

	viper.SetConfigName("envs")

	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

}
