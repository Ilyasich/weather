package rest

import (

	"net/http"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models" 
	"github.com/gin-gonic/gin"
)

func HandleCurrentWeather(r *Rest)(ctx *gin.Context) {
	lang := config.Lang

	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(400, gin.H{"error": "Требуется параметр city"})
		return
	}

	
//запрос к API погоды getCurrentWeather
	weatherData, err:= r.service.service.GetCurrentWeather(city, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить текущую погоду"})
		return
	}

	ctx.JSON(http.StatusOK, weatherData)
}

