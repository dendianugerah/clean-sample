package handler

import (
	"ddd-sample/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	CreateTask() echo.HandlerFunc
	GetTask() echo.HandlerFunc
	UpdateTask() echo.HandlerFunc
	DeleteTask() echo.HandlerFunc
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

type reqTask struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTask struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTaskUser struct {
	ID      int          `json:"id"`
	UserID  int          `json:"user_id"`
	Title   string       `json:"title"`
	Content string       `json:"content"`
	User    responseUser `json:"user"`
}

func (th *taskHandler) CreateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req reqTask
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTask, err := th.taskUsecase.Create(req.UserID, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      createdTask.ID,
			UserID:  createdTask.UserID,
			Title:   createdTask.Title,
			Content: createdTask.Content,
		}
		return c.JSON(http.StatusCreated, res)
	}
}

func (th *taskHandler) GetTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTask, err := th.taskUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      foundTask.ID,
			UserID:  foundTask.UserID,
			Title:   foundTask.Title,
			Content: foundTask.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) UpdateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req reqTask
		err = c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		updatedTask, err := th.taskUsecase.Update(id, req.UserID, req.Title, req.Content)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:      updatedTask.ID,
			UserID:  updatedTask.UserID,
			Title:   updatedTask.Title,
			Content: updatedTask.Content,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.taskUsecase.Delete(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
