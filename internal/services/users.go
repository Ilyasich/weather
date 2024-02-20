package services

import (
	"context"

	"github.com/Ilyasich/weather/internal/models"
)

// интерфейс c методами
type UsersRepository interface {
	AddUser(models.User)
	FindUser(string) bool
	GetFavorite(userToken string) ([]models.FavoriteCity, error)
	GetCurrentWeather(userToken string) models.WeatherResponse
	GetFavorites(userToken string) ([]models.FavoriteCity, error)
	SaveFavorite(userToken string, favorite models.FavoriteCity) error
	DeleteFavorite(userToken, city string) error
	SaveToken(token string, username string)
}

// Эта структура используется для предоставления сервисных функций, связанных с пользователями.
type Service struct {
	repo UsersRepository
}

// Эта функция используется для создания нового объекта `Service` с переданным репозиторием.
func New(repo UsersRepository) Service {
	return Service{
		repo: repo,
	}
}

// Метод структуры принимает контекст и пользователя `user` типа `models.User` в качестве параметров. Внутри метода вызывается метод `AddUser` чтобы добавить нового пользователя.
func (s *Service) CreateNewUser(ctx context.Context, user models.User) error {
	s.repo.AddUser(user)
	return nil
}

// Это метод `UserExists` структуры `Service`. Внутри метода вызывается метод `FindUser` репозитория `s.repo`, чтобы проверить, существует ли пользователь с указанным именем.
func (s *Service) UserExists(ctx context.Context, name string) (bool, error) {
	ok := s.repo.FindUser(name)
	return ok, nil
}

// метод получить favorit
func (s *Service) GetFavorites(ctx context.Context, userToken string) ([]models.FavoriteCity, error) {
	return s.repo.GetFavorites(userToken)
}

func (s *Service) GetCurrentWeather(ctx context.Context, userToken string) ([]models.FavoriteCity, error) {
	return s.repo.GetFavorites(userToken)
}

func (s *Service) SaveFavorite(ctx context.Context, userToken string, favorite models.FavoriteCity) error {
	return s.repo.SaveFavorite(userToken, favorite)
}

func (s *Service) DeleteFavorite(ctx context.Context, userToken, city string) error {
	return s.repo.DeleteFavorite(userToken, city)
}
