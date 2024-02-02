package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/services"
)

func CreateServer() {

	gin.SetMode(gin.ReleaseMode) // убирает лишние сообщения в консоли
	g := gin.Default()

	//тип запроса current
	g.GET("/current/:city", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "city") //??

		location := ctx.Param("city") //
		urlCurrent := fmt.Sprintf("http://api.weatherapi.com/v1/current.jso?key=%s&q=%s", url, apiKey, location)
		ctx.String(http.StatusOK, "city", location)
		resp, err := http.Get(urlCurrent)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить прогноз погоды."})
			return
		}
		defer resp.Body.Close() //закрываем калитку

		ctx.JSON(resp.StatusCode, resp.Body)
	})

	//тип запроса forecast
	g.GET("/forecast/:city", func(ctx *gin.Context) {
		location := ctx.Param("city")
		urlForecast := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s", url, apiKey, location)
		ctx.String(http.StatusOK, "city", location)

		resp, err := http.Get(urlForecast)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить прогноз погоды."})
			return
		}
		defer resp.Body.Close() //закрываем калитку

		ctx.JSON(resp.StatusCode, resp.Body)
	})

	//запуск функции клиента
	g.GET("/weather", func(ctx *gin.Context) {
		weather := users.WeatherGet() //запуск
		url := fmt.Sprintf("%weather.json?key=%s&q=%s", url, apiKey, weather)
		ctx.String(http.StatusOK, "city", weather)
		resp, err := http.Get(url)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"data": weather})
			return
		}
		defer resp.Body.Close()

		ctx.JSON(resp.StatusCode, resp.Body)
	})

}

func GetCurrentWeather(c *gin.Context) {
	city := c.Query("city")
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var weatherRes WeatherResponse
	err = json.Unmarshal(body, &weatherRes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weatherRes)
}

func GetForecastWeather() {

}

// //
func SaveFavorites(c *gin.Context) {
	login := c.Param("login")

	// Парсинг данных из тела запроса
	var bookmark Bookmark
	if err := c.ShouldBindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка наличия пользователя в списке закладок
	if bookmarks == nil {
		bookmarks = make(map[string][]Bookmark)
	}

	// Сохранение закладки для пользователя
	bookmarks[login] = append(bookmarks[login], bookmark)

	c.JSON(http.StatusOK, gin.H{"message": "Закладка успешно сохранена"})
}
