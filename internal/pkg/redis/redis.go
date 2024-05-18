package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/singhdurgesh/rednote/configs"

	"github.com/redis/go-redis/v9"
)

func Connect(config *configs.Redis) *redis.Client {

	// logger := logger.LogrusLogger

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)

	var tls_config *tls.Config

	if config.Encrypt {
		tls_config = &tls.Config{
			MinVersion: tls.VersionTLS12,
			//Certificates: []tls.Certificate{cert}
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:      fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username:  config.UserName,
		Password:  config.Password,
		DB:        config.Db,
		TLSConfig: tls_config,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	fmt.Printf(`🍟: Successfully connected to Redis at ` + address)

	return rdb

}
