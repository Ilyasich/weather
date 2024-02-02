package main

import (

	"github.com/Ilyasich/weather/internal/repostories/memory"
	"github.com/Ilyasich/weather/internal/services"
	"github.com/Ilasich/weather/internal/transport/rest"

	
)

func main() {
	repo := &memory.Repository{}
	service := services.New(repo)

	repo.NewServer(service).Run(":8080")
}
