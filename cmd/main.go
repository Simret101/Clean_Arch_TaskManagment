package main

import (
	"task/api/controller"
	"task/api/middleware"
	"task/api/router"
	"task/config"
	"task/repository"
	"task/usecase"
)

func main() {

	config.LoadConfig()

	userRepo := repository.NewInMemoryUserRepository()
	taskRepo := repository.NewInMemoryTaskRepository()

	userUsecase := usecase.NewUserUsecase(userRepo)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	authMiddleware := middleware.NewAuthMiddleware(userUsecase)

	r := router.SetupRouter(userController, taskController, authMiddleware)

	r.Run(":8080")
}
