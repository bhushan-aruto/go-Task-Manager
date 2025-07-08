package handler

import (
	"github.com/bhushan-aruto/go-task-manager/internal/entity"
	"github.com/bhushan-aruto/go-task-manager/internal/infrastructure/server/model"
	"github.com/bhushan-aruto/go-task-manager/internal/usecase"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHandler struct {
	usecase *usecase.TaskUsecase
}

func NewTaskHandler(usecase *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		usecase: usecase,
	}
}

func (h *TaskHandler) CreatTaskHandler(ctx echo.Context) error {
	var req model.CreateTaskRequest

	if err := ctx.Bind(&req); err != nil {

		return ctx.JSON(400, echo.Map{
			"error": "invalid request body",
		})
	}

	userID, ok := ctx.Get("user_id").(string)
	if !ok || userID == "" {
		return ctx.JSON(401, echo.Map{"error": "unauthorized"})
	}

	err := h.usecase.CreateTaskUsecase(userID, req.Title, req.Description)
	if err != nil {
		return ctx.JSON(500, echo.Map{"error": "could not create task"})
	}

	return ctx.JSON(201, echo.Map{
		"message": "task created successfully",
	})
}

func (h *TaskHandler) GetTaskHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	task, err := h.usecase.GetTaskByID(id)
	if err != nil {
		return ctx.JSON(404, echo.Map{
			"error": "error occured while gettong the task",
		})
	}
	return ctx.JSON(200, echo.Map{
		"data": task,
	})

}

func (h *TaskHandler) UpdatedTaskHandler(ctx echo.Context) error {
	var req model.UpdateTaskRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid body",
		})
	}

	taskID := ctx.Param("id")
	if taskID == "" {
		return ctx.JSON(400, echo.Map{"error": "task ID is required in URL"})
	}

	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return ctx.JSON(400, echo.Map{"error": "invalid task ID"})
	}

	userID := ctx.Get("user_id")
	if userID == nil {
		return ctx.JSON(401, echo.Map{"error": "unauthorized"})
	}

	task := &entity.Task{
		ID:          objID,
		UserId:      userID.(string),
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	}

	if err := h.usecase.UpdateTask(task); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{"message": "task updated successfully"})
}

func (h *TaskHandler) DeleteTaskHandler(ctx echo.Context) error {
	id := ctx.Param("id")

	err := h.usecase.DeleteTaskByID(id)
	if err != nil {
		return ctx.JSON(404, echo.Map{
			"error": "error occured while deleteing the task",
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": "task deleted succesfully",
	})
}

func (h *TaskHandler) ListTasksByUserHandler(ctx echo.Context) error {
	userId, ok := ctx.Get("user_id").(string)
	if !ok || userId == "" {
		return ctx.JSON(401, echo.Map{"error": "unauthorized"})
	}

	tasks, err := h.usecase.ListTasksByUser(userId)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, echo.Map{
		"message": "tasks fetched successfully",
		"count":   len(tasks),
		"data":    tasks,
	})
}

func (h *TaskHandler) MarkTaskAsCompleteHandler(ctx echo.Context) error {
	taskID := ctx.Param("id")

	if taskID == "" {
		return ctx.JSON(400, echo.Map{
			"error": "task ID is required",
		})
	}

	err := h.usecase.MarkTaskAsComplete(taskID)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, echo.Map{
		"message": "Task marked as complete",
	})
}
