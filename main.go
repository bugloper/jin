package main

import (
	"jin/adapters/http"
	"jin/adapters/repository"
	"jin/app"
	"jin/infra"
	"log"
)

func main() {
	userRepo := repository.NewInMemoryUserRepository()

	userService := app.NewUserService(userRepo)

	userHandler := http.NewUserHandler(userService)

	router := infra.SetupRouter(userHandler)

	log.Println("Server is running on port http://localhost:8080")
	router.Run(":8080")
}
