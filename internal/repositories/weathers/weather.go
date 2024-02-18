package main

import (
	"github.com/go-resty/resty/v2"

	
	//"github.com/Ilyasich/weather/internal/transport/rest"
	"github.com/Ilyasich/weather/internal/config"

	"fmt"
)

// ревлизуем походы во внешние API 

func GetCurrentWeather() {
	client := resty.New()

	resp, err := client.R().
	SetQueryParams(map[string]string{
		"q":    config.City,
		"key":  config.Apikey,
	}).
	Get("https://api.weatherapi.com/v1/current.json")

	fmt.Println(resp, err)

	
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


