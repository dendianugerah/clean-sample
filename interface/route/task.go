package route

import (
	"ddd-sample/interface/handler"

	"github.com/labstack/echo/v4"
)

func InitTaskRoute(e *echo.Echo, taskHandler handler.TaskHandler) {
	e.POST("/task", taskHandler.CreateTask())
	e.GET("/task/:id", taskHandler.GetTask())
	e.PUT("/task/:id", taskHandler.UpdateTask())
	e.DELETE("/task/:id", taskHandler.DeleteTask())
}
