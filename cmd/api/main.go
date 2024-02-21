package main

import (
	"github.com/Ilyasich/weather/internal/services"
	"github.com/Ilyasich/weather/internal/transport/rest"
	//"github.com/Ilyasich/weather/internal/transport/memory"
)


func main() {

	//rest.NewServer(*services.Service).Run(":8080")
	//repo := &memory.Repository{}
	//service := services.New(repo)

	
	server := rest.NewServer(&services.Service{})
	server.Run(":8080")
   }