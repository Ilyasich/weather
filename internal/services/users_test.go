package services

import (
	//"context"
	"testing"

	// "github.com/Ilyasich/weather/internal/services"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// //библиотека gomok генерирует временные инплементации под интерфейсы

func TestDeleteFavorite(t *testing.T) {
	name := "dima"
	city := "batumi"

	//fake repo

	ctrl := gomock.NewController(t)
	repo := NewMockUserRepository(ctrl)
	repo.EXPECT().FindUser(name).Return(false)//ожидай что придет запрос FindUser и на него верни false
		svc := Service{repo}

	ok := svc.DeleteUserFavorite(name, city)
	require.False(t, ok)
}



// func TestDeleteFavorite(t *testing.T) {
// 	name := "dima"
// 	city := "batumi"

// 	repo := services.New()
// 	svc := Service{repo}

// 	ok := svc.DeleteFavorite(context.Context, city, name)
// 	require.False(t, ok)

// }
