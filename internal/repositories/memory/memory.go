package memory

import (
	// "context"
	// "fmt"

	"github.com/Ilyasich/weather/internal/models"
	//"github.com/Ilyasich/weather/internal/pkg/tracing"
)

var favorites = make(map[string]models.FavoriteCity)


// будет использоваться для хранения пользователей.
type Repository struct {
	users []models.User
}

// DeleteFavorite implements services.UsersRepository.
func (*Repository) DeleteFavorite(userToken string, city string) error {
	panic("unimplemented")
}

// GetCurrentWeather implements services.UsersRepository.
func (*Repository) GetCurrentWeather(userToken string) models.WeatherResponse {
	panic("unimplemented")
}

// GetFavorite implements services.UsersRepository.
func (*Repository) GetFavorite(userToken string) ([]models.FavoriteCity, error) {
	panic("unimplemented")
}

// GetFavorites implements services.UsersRepository.
func (*Repository) GetFavorites(userToken string) ([]models.FavoriteCity, error) {
	panic("unimplemented")
}

// GetUserToken implements services.UsersRepository.
func (*Repository) GetUserToken(token string) (string, bool) {
	panic("unimplemented")
}

// SaveFavorite implements services.UsersRepository.
func (*Repository) SaveFavorite(userToken string, favorite models.FavoriteCity) error {
	panic("unimplemented")
}

// SaveToken implements services.UsersRepository.
func (*Repository) SaveToken(token string, username string) {
	panic("unimplemented")
}

// для добавления нового пользователя в `r.users`.
func (r *Repository) AddUser(name models.User) {
	r.users = append(r.users, name)
}

// Это метод `FindUser` определенный для структуры `Repository`принимает аргумент name
func (r *Repository) FindUser(name string) bool {
	for _, u := range r.users { //проверка пользователя на существование
		if u.Name == name {
			return true
		}
	}
	return false
}
