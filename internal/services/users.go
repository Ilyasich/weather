package services

import (
	"context"

	"github.com/Ilyasich/weather/internal/models"
)


//ручки

type UserRepository interface {
	AddUser(models.User)
	FindUser(string) bool
}

type Service struct {
	repo UserRepository
}

func New(repo UserRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) CreateNewUser(ctx context.Context, user models.User) error {
	s.repo.AddUser(user)

	return nil
}

func (s *Service) UserExists(ctx context.Context, name string) (bool, error) {
	ok := s.repo.FindUser(name)
	return ok, nil
}

//получить favorite
func (s *Service) GetFavorites(ctx context.Context, userToken string) ([]models.Favorite, error) {
	return s.repo.GetFavorites(userToken)

}

//сохранить favorit 
func (s *Service) SaveFavorite(ctx context.Context, userToken string, favorite models.Favorite) error {
	return s.repo.SaveFavorite(userToken, favorite)
}

//удалить favorit
func (s *Service) DeleteFavorite(ctx context.Context, userToken, city string) error {
	return s.repo.DeleteFavorite(userToken, city)
}
