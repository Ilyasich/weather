package memory

import (
	// "context"
	// "fmt"

	"errors"

	"github.com/Ilyasich/weather/internal/models"
	//"github.com/Ilyasich/weather/internal/pkg/tracing"
)



// будет использоваться для хранения пользователей.
type Repository struct {
	users []models.User
	favoritMap map[string][]models.FavoriteCity
	tokenMap map[string]string
}



// DeleteFavorite implements services.UsersRepository.
func (r *Repository) DeleteFavorite(userToken string, city string) error {
	favorit, exist := r.favoritMap[userToken]
	if !exist {
		return errors.New("not found favorit")
	}

	for i, fav := range favorit {
		if fav.City == city {
			r.favoritMap[userToken] = append(favorit[:i], favorit[i+1:]... )
			return nil
		}
	}

	return errors.New("not found favorit")
}




// GetFavorites 
func (r *Repository) GetFavorite(userToken string) ([]models.FavoriteCity, error) {
	return r.favoritMap[userToken], nil
}

// GetUserToken implements services.UsersRepository.???
func (*Repository) GetUserToken(token string) (string, bool) {
	panic("unimplemented")
}

// SaveFavorite .
func (r *Repository) SaveFavorite(userToken string, favorite models.FavoriteCity) error {
	r.favoritMap[userToken] = append(r.favoritMap[userToken], favorite)
	return nil
}

// SaveToken implements services.UsersRepository.
func (r *Repository) SaveToken(token string, username string) error {
	r.tokenMap[token] = username
	return nil
}

// для добавления нового пользователя в `r.users`.
func (r *Repository) AddUser(name models.User) bool {
	r.users = append(r.users, name)
	return true
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


