package main

import (
	"go-rest-api-handson/controller"
	"go-rest-api-handson/db"
	"go-rest-api-handson/repository"
	"go-rest-api-handson/router"
	"go-rest-api-handson/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
