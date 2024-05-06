package main

import (
	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/app/services"
	"github.com/singhdurgesh/rednote/internal/jobs"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	"github.com/singhdurgesh/rednote/internal/pkg/postgres"
)

func main() {
	configs.LoadConfig() // Setup Configuration

	logger.Init()

	// connect Database
	postgres.Connect(&configs.EnvConfig.Postgres)

	services.Init()

	jobs.ConsumerStart()
}
