package services

import (
	"context"

	"github.com/Ilyasich/weather/internal/models"
)

type UsersRepository interface {//интерфейс с двумя методами
	AddUser(models.User)//метод
	FindUser(string) bool//метод
}

//Эта структура используется для предоставления сервисных функций, связанных с пользователями.
type Service struct {
	repo UsersRepository
}

//Эта функция используется для создания нового объекта `Service` с переданным репозиторием.
func New(repo UsersRepository) Service {
	return Service{
		repo: repo,
	}
}

//Метод структуры принимает контекст и пользователя `user` типа `models.User` в качестве параметров. Внутри метода вызывается метод `AddUser` чтобы добавить нового пользователя.
func (s *Service) CreateNewUser(ctx context.Context, user models.User) error {
	s.repo.AddUser(user)

	return nil
}

// Это метод `UserExists` структуры `Service`. Внутри метода вызывается метод `FindUser` репозитория `s.repo`, чтобы проверить, существует ли пользователь с указанным именем.
func (s *Service) UserExists(ctx context.Context, name string) (bool, error) {
	ok := s.repo.FindUser(name)
	return ok, nil
}



// //??? метод сохранить favorit
func (s *Service) GetFavorites(ctx context.Context, userToken string) (error) {
	return s.repo.GetFavorites(userToken)
}