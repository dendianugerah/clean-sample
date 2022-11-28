package main

import (
	"ddd-sample/config"
	infra "ddd-sample/infra"
	"ddd-sample/interface/handler"
	"ddd-sample/interface/route"
	"ddd-sample/usecase"

	_ "ddd-sample/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()
	db := config.NewDB()

	//user
	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	route.InitUserRoute(e, userHandler)

	//task
	taskRepository := infra.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)
	route.InitTaskRoute(e, taskHandler)

	//swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
