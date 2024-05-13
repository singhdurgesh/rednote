package rednote

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/app/services"
	"github.com/singhdurgesh/rednote/internal/jobs/task_register"
	"github.com/singhdurgesh/rednote/internal/jobs/task_server"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	"github.com/singhdurgesh/rednote/internal/pkg/postgres"
	"github.com/singhdurgesh/rednote/internal/router"
)

func Init() {
	// init router
	router.Init()
	r := router.Router

	configs.LoadConfig() // Setup Configuration
	EnvConfig := configs.EnvConfig

	logger.Init()
	logger := logger.LogrusLogger

	// connect Database
	postgres.Connect(&configs.EnvConfig.Postgres)

	// Start Work Task Server
	task_server.StartServer()
	task_register.RegisterTasks()

	// go jobs.ConsumerStart()
	// Service Initialization
	services.Init()
	// graceful shutdown
	server := &http.Server{
		Addr:    EnvConfig.Server.Port,
		Handler: r,
	}

	logger.Printf("👻 Server is now listening at port:  %s\n", EnvConfig.Server.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	logger.Println("server receive a signal: ", i.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("server shutdown error: %s\n", err)
	}
	logger.Println("Server exiting")
}
