package rest

import (

	"net/http"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"
	"github.com/gin-gonic/gin"
)

func handleCurrentWeather(r *Rest)(ctx *gin.Context) {
	lang := config.Lang

	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(400, gin.H{"error": "Требуется параметр city"})
		return
	}

	

	weatherData, err:= GetCurrentWeather(city, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить текущую погоду"})
		return
	}

	ctx.JSON(http.StatusOK, weatherData)
}

