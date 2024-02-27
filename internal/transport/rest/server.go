package rest

import (
	"io"
	"net/http"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Rest struct {
	lg *zap.SugaredLogger
	service *services.Service
}

func NewServer(host string, service *services.Service) *gin.Engine {
	if config.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = io.Discard
	g := gin.Default()


	rest := Rest{lg, service}
	

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	g.Use(cors.New(config))
	g.Use(func(ctx *gin.Context) {
		lg.Info("http request", ctx.Request.URL.Path)
	})

	g.GET("/users/:name/exists", rest.userExists) //проверка существования пользователя
	g.POST("/users", rest.createUser)
	//g.POST("/login", rest.login)



	
	g.GET("/weather/current", rest.handleCurrentWeather) //получение текущая погода
	g.POST("/favorites", rest.createFavorite)
	g.GET("/favorites", rest.getFavorites)
	g.DELETE("/favorites/:city", rest.deleteFavorite)

	

	return &http.Server{
		Addr: cfg.ServerHost,
		Handler: r,
	}

	return g
}
