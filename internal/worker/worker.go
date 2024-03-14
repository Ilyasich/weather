package worker

import (
	"fmt"
	"time"
)



type WeatherService struct{}

func (ws *WeatherService) CheckWeather(city City) (string, error) {
	// проверкa погоды здесь
	return "Rain", nil
}

type UserService struct {
	Users []User
}

func (us *UserService) GetUserAlerts(userID int) ([]Alert, error) {
	// получения списка алертов пользователя здесь
	return []Alert{}, nil
}


	// Инициализация сервисов
	ws := &WeatherService{}
	us := &UserService{}

	// Инициализация worker
	go func() {
		for {
			for _, user := range us.Users {
				for _, city := range user.Favorites {
					weather, err := ws.CheckWeather(city)
					if err != nil {
						fmt.Println(err)
						continue
					}

					if weather == "Rain" {
						user.Alerts = append(user.Alerts, Alert{
							City:  city,
							Alert: "It's going to rain today!",
						})
					}
				}
			}

			// Проверка раз в сутки
			time.Sleep(24 * time.Hour)
		}
	}()


