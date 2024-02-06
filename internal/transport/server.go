package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/services"
)

type Rest struct {
	service *services.Service
}

func NewServer(service services.Service) *gin.Engine {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	rest := Rest{&services.Service{}}

	//тут дергаем ручки
	r.GET("/users/:name/exists", rest.userExists)
	return r

	r.POST("/favorites", rest.CreateFovorite)
}
