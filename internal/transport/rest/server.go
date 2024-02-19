package rest

import (
	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/services"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	service *services.Service
}

func NewServer(service *services.Service) *gin.Engine {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.Default()

	rest := Rest{service}

	g.GET("/users/:name/exists", rest.userExists) //проверка существования пользователя
	//g.POST("/users", rest.CreateNewUser)

	g.GET("/weather/current", rest.handleCurrentWeather) //получение текущая погода
	g.POST("/favorites", rest.createFavorite)
	g.GET("/favorites", rest.getFavorites)
	//g.DELETE("/favorites/:city", rest.deleteFavorite)

	return g

}
