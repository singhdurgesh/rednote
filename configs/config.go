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

func LoadConfig() *Config {
	viper.SetEnvPrefix("rednote")
	viper.SetDefault("env", "local") // Defauld ENV Set to Local
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()

	path, err := os.Getwd() // get curent path
	if err != nil {
		panic(err)
	}

	env := viper.Get("env")
	fmt.Println(env)
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/configs/environments")

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

	return config
}
