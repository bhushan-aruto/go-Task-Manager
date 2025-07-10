package server

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/bhushan-aruto/go-task-manager/config"
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/database"
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/server/handler"
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/server/routes"
	"github.com/bhushan-aruto/go-task-manager/internal/usecase"
)

func Start(cfg *config.Config) {
	db := database.ConnecttoMOngo(cfg)

	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	userRepo := database.NewUserRepo(db)
	taskRepo := database.NewTaskRepo(db)

	userUsecase := usecase.NewUserUsecase(userRepo)
	taskUsecase := usecase.NewTaskUseCaseRepo(taskRepo)

	userHandler := handler.NewUserHandler(userUsecase)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	routes.InitRoutes(e, userHandler, taskHandler)

	if err := e.Start(cfg.ServerAdres); err != nil {
		log.Fatal("Server failed to start: ", err)
	}

}
