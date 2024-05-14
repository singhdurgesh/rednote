package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres   Postgres
	Redis      Redis
	Server     Server
	Jwt        Jwt
	AMQPConfig AMQPConfig
	App        App
}

var EnvConfig *Config

func LoadConfig() *Config {
	viper.AllowEmptyEnv(true)
	viper.BindEnv("go_env")

	path, err := os.Getwd() // get curent path
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/configs")

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		fmt.Println("hellss")
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Fetch Configuration from ENV File
	viper.SetConfigType("env")
	viper.SetConfigFile((".env"))
	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		fmt.Println("hellss")
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	EnvConfig = config

	return config
}
