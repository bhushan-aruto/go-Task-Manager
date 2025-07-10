package routes

import (
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/server/handler"
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/server/middlewares"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, userHandler *handler.UserHandler, taskHandler *handler.TaskHandler) {

	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)

	//auther
	authGroup := e.Group("/api", middlewares.JWTMiddleware)

	authGroup.POST("/tasks", taskHandler.CreatTaskHandler)
	authGroup.GET("/tasks", taskHandler.ListTasksByUserHandler)
	authGroup.GET("/tasks/:id", taskHandler.GetTaskHandler)
	authGroup.PUT("/tasks/:id", taskHandler.UpdatedTaskHandler)
	authGroup.DELETE("/tasks/:id", taskHandler.DeleteTaskHandler)
	authGroup.PUT("/tasks/:id/completed", taskHandler.MarkTaskAsCompleteHandler)

}
