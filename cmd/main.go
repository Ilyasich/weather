package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ilyasich/weather/internal/config"
	
	"github.com/Ilyasich/weather/internal/repostories/memory"
	"github.com/Ilyasich/weather/internal/services"
	"github.com/Ilyasich/weather/internal/transport/rest"
	"go.uber.org/zap"
)

func main() {
	//
	logCfg := zap.NewDevelopmentConfig()
	logCfg.OutputPaths = []string{"server.log"}
	logCfg.Encoding = "json"

	logger, _ := logCfg.Build()
	defer logger.Sync()
	lg := logger.Sugar()//список интерфейсов с логами

	//конфигурация сервера
	cfg, err := config.Read()
	if err != nil {
		lg.Fatal(err)
	}

	repo := &memory.Repository{}
	service := services.New(repo)

	server := rest.NewServer(lg, cfg.Server, service)

	//мягкая остановка сервера
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			lg.Panicln("Shutdown error:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		lg.Panicln(err)
	}
}
