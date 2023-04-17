package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
	"todo-app/internal/config"
	"todo-app/internal/todo"
	"todo-app/internal/todo/db"
	"todo-app/pkg/client/mongodb"
	"todo-app/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Router created")
	router := httprouter.New()

	cfg := config.GetConfig()
	cfgMongo := cfg.Database

	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.User, cfgMongo.Password, cfgMongo.Name, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongoDBClient, cfg.Database.Collection, logger)

	todo1 := todo.Todo{
		ID:        "1",
		Title:     "asdasdasd",
		Completed: false,
	}
	todo1ID, err := storage.Create(context.Background(), todo1)
	if err != nil {
		panic(err)
	}
	logger.Info(todo1ID)

	logger.Info("Registering todo handler")
	handler := todo.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Starting application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("Detecting app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Creating socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("Listening unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("Server listening on unix socket: %s", socketPath)
	} else {
		logger.Info("Listening tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("Server listening on %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
