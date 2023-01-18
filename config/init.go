package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Error on load config: %s \n", err.Error()))
	}

}
