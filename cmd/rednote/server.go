package rednote

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/pkg/amqp"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	"github.com/singhdurgesh/rednote/internal/pkg/postgres"
	"github.com/singhdurgesh/rednote/internal/pkg/redis"
	"github.com/singhdurgesh/rednote/internal/router"
	"github.com/singhdurgesh/rednote/internal/tasks/task_register"
)

func Init() {
	app.Config = configs.LoadConfig() // Setup Configuration

	// Connect Logger
	app.Logger = logger.Init()

	// connect Database
	app.Db = postgres.Connect(&app.Config.Postgres)

	// Connect Cache
	app.Cache = redis.Connect(&app.Config.Redis)

	// Start Work Task Server
	app.Broker = amqp.Connect(&app.Config.AMQPConfig)
	task_register.RegisterTasks()

	// init router
	router.Init()
	r := router.Router
	// graceful shutdown
	server := &http.Server{
		Addr:    app.Config.Server.Port,
		Handler: r,
	}

	app.Logger.Printf("👻 Server is now listening at port:  %s\n", app.Config.Server.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatalf("server listen error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	i := <-quit
	app.Logger.Println("server receive a signal: ", i.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		app.Logger.Fatalf("server shutdown error: %s\n", err)
	}
	app.Logger.Println("Server exiting")
}
