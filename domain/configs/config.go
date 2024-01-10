package configs

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func ReadConfig(config *Configuration) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(pwd + "/../configs")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

}
