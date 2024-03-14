package worker

import "time"

type City struct {
	Name string `json:"name"`
}

type User struct {
	Favorites []City
	Alerts    []Alert
}

type Alert struct {
	City  City
	Date time.Time
}