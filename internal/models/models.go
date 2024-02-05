package models

type Favorite struct {
	City       string   `json:"city"`
	Paramertrs []string `json:"parametrs"`
}

type CurrentWeatherResponse struct {
	Location struct {
		Name      string  `json:"name"`
		Region    string  `json:"region"`
		Country   string  `json:"country"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		TzID      string  `json:"tz_id"`
		Localtime string  `json:"localtime"`
		Epoch     int64   `json:"epoch"`
		UtcOffset float64 `json:"utc_offset"`
	} `json:"location"`
	Current struct {
		LastUpdated string  `json:"last_updated"`
		TempC       float64 `json:"temp_c"`
		Condition   struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		Wind struct {
			Kph float64 `json:"kph"`
			Dir string  `json:"dir"`
		} `json:"wind"`
	} `json:"current"`
}
