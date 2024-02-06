package rest

import (
	"context"
	"net/http"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"
	"github.com/gin-gonic/gin"
)

//тут пишем все функции


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

//?
func (s *Rest) userExists(ctx *gin.Context) {
	ok, err := s.service.UserExists(ctx, ctx.Param("name"))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Status bool
	}{
		Status: ok,
	})
}

func (s *Rest) CreateFovorite(ctx *gin.Context) {
	token := ctx.Param("token")
	var data map[string]string
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Недопустимый запрос"})
		return
	}
if err := s.service.SaveFavorite(ctx, token, models.Favorite{}); err != nil {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Фаил не сохранен"})
	return
}

ctx.JSON(http.StatusOK, gin.H{"message":"Успешно сохранено"})

}



func HandleCurrentWeather(r *Rest) (ctx *gin.Context) {
	lang := config.Lang
	token := ctx.Query("token")

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
