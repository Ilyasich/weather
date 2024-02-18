package rest

import (
	"fmt"
	"net/http"
	"os/user"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"
	"github.com/gin-gonic/gin"
)

//Метод `createUser` обрабатывает POST-запрос на создание нового пользователя.
func (s *Rest) createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)//преобразует JSON в объект модели `User` 
	if err != nil {
		s.lg.Error("Invalid body")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = s.service.CreateNewUser(ctx, user)//создает нового пользователя в хранилище данных.
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

//Метод `userExists` обрабатывает GET-запрос на проверку существования пользователя по имени.
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


//сохранение закладок для пользователя
func (q *Rest) SaveFavorites(ctx *gin.Context) {
	
	login := ctx.Param("login")
	var favoriteReq models.Favorite
	err := ctx.ShouldBindJSON(&favoriteReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f := models.Favorite{
		City: favoriteReq.City,
		Parameters: favoriteReq.Parameters,
	}

	models.Favorites[login] = append(models.Favorites[login], f)
	
}

func (q *Rest) GetFavorites(ctx *gin.Context) {
	login := ctx.Param("login")
	favorites, ok := models.Favorites[login]
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"favorites": models.Favorite{}})
	}
	ctx.JSON(http.StatusOK, favorites)
}

func (q *Rest) handleCurrentWeather(ctx *gin.Context) {
	//city := ctx.Query("city")
	
	url := fmt.Sprintf("%weather.json?key=%s&q=%s", config.City, config.Apikey)
	ctx.String(http.StatusOK, "city")
	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"data": url})
		return
	}
	defer resp.Body.Close()

	city := ctx.Query("city")
		if city == "" {
			favorites, err := q.service.GetFavorites()
			if err != nil {
				city = config.City
			}

			
		}
		// Здесь вы можете добавить логику для получения погоды для города.
		ctx.JSON(http.StatusOK, gin.H{"city": city, "weather": "Sunny"})
	
	

	ctx.JSON(resp.StatusCode, resp.Body)
}



