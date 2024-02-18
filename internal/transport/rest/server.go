package rest

import (
	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/services"
	
	"github.com/gin-gonic/gin"

	//"github.com/Ilyasich/weather/internal/rest"

	"fmt"
	"net/http"
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

	g.GET("/weather/current", func(ctx *gin.Context) {
		url := fmt.Sprintf("%weather.json?key=%s&q=%s", config.City, config.Apikey)
		ctx.String(http.StatusOK, "city")
		resp, err := http.Get(url)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"city": url})
			return
		}
		defer resp.Body.Close()

		ctx.JSON(resp.StatusCode, resp.Body)
	})


	g.GET("/users/:name/exists", rest.UserExists)//проверка существования пользователя
	g.POST("/users", rest.createUser)
	


	g.GET("/weather/current", rest.handleCurrentWeather)//получение текущая погода
	g.POST("/favorites", rest.createFavorite)
	g.GET("/favorites", rest.getFavorites)
	g.DELETE("/favorites/:city", rest.deleteFavorite)


return g

}

