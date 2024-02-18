package rest

import (
	//"fmt"
	"net/http"
	//"os/user"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"
	"github.com/gin-gonic/gin"
)

//Метод `createUser` обрабатывает POST-запрос на создание нового пользователя.
func (s *Rest) CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)//преобразует JSON в объект модели `User` 
	if err != nil {
		//s.lg.Error("Invalid body")
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

// func (q *Rest) createFavorite(ctx *gin.Context) {
// 	var fav models.Favorite
// 		if err := ctx.ShouldBindJSON(&fav); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		login := ctx.Param("login")
// 		models.Favorites[login] = fav
// 		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
// 	}
	

func (q *Rest) handleCurrentWeather(ctx *gin.Context) {
	city := ctx.Query("city")


	if city == "" {
		// Если город не указан, получаем список избранных городов пользователя
		favorites, err := q.service.GetFavorites(q, models.User)
		if err != nil || len(favorites) == 0 {
			// Если у пользователя нет избранных городов, используем город по умолчанию
			city = config.DefoultCity
		} else {
			// Используем город из первой закладки
			city = favorites[0].City
		}
	}


	weatherData, err := q.service.GetCurrentWeather(city, config.Lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current weather"})
		return
	}

	// Отправляем полученные данные о погоде клиенту
	ctx.JSON(http.StatusOK, weatherData)
}
	
	// url := fmt.Sprintf("weather.json?key=%s&q=%s", config.City, config.Apikey)
	// ctx.String(http.StatusOK, "city")
	// resp, err := http.Get(url)
	// if err != nil {
	// 	ctx.JSON(http.StatusOK, gin.H{"data": url})
	// 	return
	// }
	// defer resp.Body.Close()
	
	
	// ctx.JSON(resp.StatusCode, resp.Body)







