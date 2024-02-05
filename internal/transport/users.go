package rest

import (
	"net/http"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"
	"github.com/gin-gonic/gin"
)


//тут пишем асе функции


func (s *Rest) createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = s.service.CreateNewUser(ctx, user)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (s *Rest) userExists(ctx *gin.Context) {
	ok, err := s.service.UserExists(ctx, ctx.Param("name"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return

	}
}

func HandleCurrentWeather(r *Rest) (ctx *gin.Context) {
	lang := config.Lang

	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(400, gin.H{"error": "Требуется параметр city"})
		return
	}

	//запрос к API погоды getCurrentWeather
	weatherData, err := r.service.GetCurrentWeather(city, lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить текущую погоду"})
		return
	}

	ctx.JSON(http.StatusOK, weatherData)
}
