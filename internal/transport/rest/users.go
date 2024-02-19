package rest

import (
	"net/http"

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


//метод сохранение закладок для пользователя
func (q *Rest) SaveFavorites(ctx *gin.Context) {
	
	login := ctx.Param("login")
	var favoriteReq models.FavoriteCity
	err := ctx.ShouldBindJSON(&favoriteReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f := models.FavoriteCity{
		City: favoriteReq.City,
		Parameters: favoriteReq.Parameters,
	}

	models.Favorites[login] = append(models.Favorites[login], f)
	
}

func (q *Rest) GetFavorites(ctx *gin.Context) {
	login := ctx.Param("login")
	favorites, ok := models.Favorites[login]
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"favorites": models.FavoriteCity{}})
	}
	ctx.JSON(http.StatusOK, favorites)
}


func (r *Rest) createFavorite(c *gin.Context) {

	username, ok := GetUserFromContext(c)
	if !ok {
		return
	}

	var favorite models.FavoriteCity
	if err := c.BindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Использование извлеченного имени пользователя для сохранения избранного
	if err := r.services.SaveFavorite(c, username, favorite); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorite saved successfully"})
}
	

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
	

func (q *Rest) getFavorites(ctx *gin.Context) {
	username, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}
	favorites, err := q.service.GetFavorites(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get favorites"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"favorites": favorites})
}


func GetUserFromContext(c *gin.Context) (string, bool) {
	usernameInterface, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in context"})
		return "", false
	}
	username, ok := usernameInterface.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username in context is not a string"})
		return "", false
	}
	return username, true
}
