package main

import (
	"github.com/Ilyasich/weather/internal/services"
	"github.com/Ilyasich/weather/internal/transport/rest"
)

func main() {

	//rest.NewServer(*services.Service).Run(":8080")
	server := rest.NewServer(service)
	server.Run(":8080")
   }