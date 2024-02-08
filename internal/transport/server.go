package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/services"
)

type Rest struct {
	service *services.Service
	lg *zap.SugaredLogger
}

func NewServer(lg *zap.SugaredLogger, cfg config.ServerConfig, services.Service) *gin.Engine {
	if cfg.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = io.Discard
	r := gin.Default()
	rest := Rest{lg, &services.Service{}}

	r.Use(func(ctx *gin.Context){
		lg.Info("http reguest", ctx.Request.URL.Path)
	})

	//тут дергаем ручки

	r.POST("/users", rest.CreateUser)
	r.GET("/users/:name/exists", rest.UserExists)

	return &http.Server{
		Addr: cfg.ServerHost,
		Handler: r,
	}

}
	

// 	r.POST("/favorites", func(ctx *gin.Context){
// 		ctx.JSON(http.StatusOK, gin.H{"message": "Успешно"})
// 	})
	
// }
