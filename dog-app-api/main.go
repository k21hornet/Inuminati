package main

import (
	"dog-app-api/controller"
	"dog-app-api/db"
	"dog-app-api/repository"
	"dog-app-api/router"
	"dog-app-api/usecase"
	"dog-app-api/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	dogValidator := validator.NewDogValidator()
	userRepository := repository.NewUserRepository(db)
	dogRepository := repository.NewDogRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	dogUsecase := usecase.NewDogUsecase(dogRepository, dogValidator)
	userController := controller.NewUserController(userUsecase)
	dogController := controller.NewDogController(dogUsecase)
	e := router.NewRouter(userController, dogController)
	//サーバーを起動
	e.Logger.Fatal(e.Start(":80"))
}
