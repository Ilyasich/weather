package rest

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"

	"github.com/gin-gonic/gin"
)

// Метод `createUser` обрабатывает POST-запрос на создание нового пользователя.
func (g *Rest) createUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user) //преобразует JSON в объект модели `User`
	if err != nil {
		//s.lg.Error("Invalid body")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err = g.service.CreateNewUser(ctx, user) //создает нового пользователя в хранилище данных.
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

// Метод `userExists` обрабатывает GET-запрос на проверку существования пользователя по имени.
func (g *Rest) userExists(ctx *gin.Context) {
	ok, err := g.service.UserExists(ctx, ctx.Param("name"))
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

// метод сохранение закладок для пользователя
func (g *Rest) SaveFavorites(ctx *gin.Context) {

	login := ctx.Param("login")
	var favoriteReq models.FavoriteCity
	err := ctx.ShouldBindJSON(&favoriteReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f := models.FavoriteCity{
		City:       favoriteReq.City,
		Parameters: favoriteReq.Parameters,
	}

	models.Favorites[login] = append(models.Favorites[login], f)

}

func (g *Rest) GetFavorites(ctx *gin.Context) {
	login := ctx.Param("login")
	favorites, ok := models.Favorites[login]
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"favorites": models.FavoriteCity{}})
	}
	ctx.JSON(http.StatusOK, favorites)
}

func (g *Rest) createFavorite(ctx *gin.Context) {

	username, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}

	var favorite models.FavoriteCity
	if err := ctx.BindJSON(&favorite); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Использование извлеченного имени пользователя для сохранения избранного
	if err := g.service.SaveFavorite(ctx, username, favorite); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save favorite"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Favorite saved successfully"})
}

func (g *Rest) handleCurrentWeather(ctx *gin.Context) {

	username, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}

	city := ctx.Query("city")

	if city == "" {
		// Если город не указан, получаем список избранных городов пользователя
		favorites, err := g.service.GetFavorites(ctx, username)
		if err != nil || len(favorites) == 0 {
			// Если у пользователя нет избранных городов, используем город по умолчанию
			city = config.DefoultCity
		} else {
			// Используем город из первой закладки
			city = favorites[0].City
		}
	}

	weatherData, err := g.service.GetCurrentWeather(ctx, config.Lang)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get current weather"})
		return
	}

	// Отправляем полученные данные о погоде клиенту
	ctx.JSON(http.StatusOK, weatherData)
}

func (g *Rest) getFavorites(ctx *gin.Context) {
	username, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}
	favorites, err := g.service.GetFavorites(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get favorites"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"favorites": favorites})
}

func GetUserFromContext(ctx *gin.Context) (string, bool) {
	usernameInterface, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in context"})
		return "", false
	}
	username, ok := usernameInterface.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Username in context is not a string"})
		return "", false
	}
	return username, true
}

func (g *Rest) deleteFavorite(ctx *gin.Context) {
	username, ok := GetUserFromContext(ctx)
	if !ok {
		return
	}

	city := ctx.Param("city")
	if err := g.service.DeleteFavorite(ctx, username, city); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete favorite"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Favorite deleted successfully"})
}


func generateTokenUser(ctx *gin.Context, login string) string {
	userJSON, _ := json.Marshal(map[string]string{"login": login})
	return base64.StdEncoding.EncodeToString(userJSON)
}


 func(g *Rest) login(ctx *gin.Context) {
	var user models.LoginRequest
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := generateTokenUser(models.Login)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}