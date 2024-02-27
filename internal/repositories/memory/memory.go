package memory

import (
	"context"
	"fmt"

	"github.com/Ilyasich/weather/internal/models"
	"github.com/Ilyasich/weather/internal/pkg/tracing"

)

var favorites = make(map[string][]models.FavoriteCity)

//будет использоваться для хранения пользователей.
type Repository struct {
	users []models.User
}

//для добавления нового пользователя в `r.users`.
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