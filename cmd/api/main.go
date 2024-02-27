package main

import (

	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ilyasich/weather/internal/services"
	"github.com/Ilyasich/weather/internal/transport/rest"
	"github.com/Ilyasich/weather/internal/transport/memory"
	"github.com/Ilyasich/weather/internal/config"
	
	"go.uber.org/zap"
)


func main() {

	logCfg := zap.NewDevelopmentConfig()
	logCfg.OutputPaths = []string{"server.log"}
	logCfg.Encoding = "json"

	logger, _ := logCfg.Build()
	defer logger.Sync()
	lg := logger.Sugar()

	cfg, err := config.Read()
	if err != nil {
		lg.Fatal(err)
	}

	repo := &memory.Repository{}
	service := services.New(repo)

	server := rest.NewServer(lg, cfg.Server, service)

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



	//rest.NewServer(*services.Service).Run(":8080")
	//repo := &memory.Repository{}
	//service := services.New(repo)

	
// 	server := rest.NewServer(&services.Service{})
// 	server.Run(":8080")
//    }