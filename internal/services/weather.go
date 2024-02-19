package services

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"
)

// ревлизуем походы во внешние API

// принимает город и язык возвращает ошибку и структуру в моделях
func GetCurrentWeather(city, lang string) (*models.WeatherResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":    "Batumi",
			"key":  config.Api_key,
			"lang": lang,
		}).
		Get("https://api.weatherapi.com/v1/current.json")

	if err != nil {
		return nil, err
	}
	//разджейсониваем структуру погоды
	var weatherResponse models.WeatherResponse
	err = json.Unmarshal(resp.Body(), &weatherResponse)
	if err != nil {
		return nil, err
	}
	return &weatherResponse, nil

}

// const(
// 	apiKey = "3caf85347f7e49e481d110120241401"
// )

// func main() {
// 	client := resty.New()

// 	resp, err := client.R().
// 	SetQueryParams(map[string]string{
// 		"key": apiKey,
// 		"q": "Batumi",
// 	}).
// 	Get("https://api.weatherapi.com/v1/current.json") //дергаем ручку за путь

// 	 fmt.Println(resp, err)
// }
