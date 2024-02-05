package services

import (
	"encoding/json"
	"fmt"

	"github.com/Ilyasich/weather/internal/config"
	"github.com/Ilyasich/weather/internal/models"
	"github.com/go-resty/resty/v2"
)


//клиент
//делает запрос к API погодыи возвращает погодные данные указанного города
func (s *Service) GetCurrentWeather(city,lang string ) (*models.CurrentWeatherResponse, error) {
	
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"key": config.ApiKey,
			"q": city,
		}).
		Get("http://api.weatherapi.com/v1/current.json")

		if err != nil {
			fmt.Println("Не удалось получить текущую погоду.")
			return nil, err
		}
		
		var weatherResponse models.CurrentWeatherResponse
		err = json.Unmarshal(resp.Body(), &weatherResponse)
		if err != nil {

			fmt.Println("Error")
			return nil, err
		}

		return &weatherResponse, nil
}