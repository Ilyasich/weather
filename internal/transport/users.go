package rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	url    = "https://api.weatherapi.com/v1/"
	apiKey = "3caf85347f7e49e481d110120241401"
)

func WeatherGet() string {

	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"key": apiKey,
			"q":   "Batumi",
		}).
		Get(url + "http://api.weatherapi.com/v1/current.json")

	if err != nil {
		return fmt.Sprintf("Не удалось узнать текущую погоду: %v", err)
	}

	return resp.String()
}
