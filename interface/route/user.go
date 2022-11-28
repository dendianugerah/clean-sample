package route

import (
	"ddd-sample/interface/handler"

	"github.com/labstack/echo/v4"
)

func InitUserRoute(e *echo.Echo, userHandler handler.UserHandler) {
	e.POST("/user", userHandler.CreateUser())
	e.GET("/user/:id", userHandler.GetUser())
	e.GET("/users", userHandler.GetAll())
	e.PUT("/user/:id", userHandler.UpdateUser())
	e.DELETE("/user/:id", userHandler.DeleteUser())
}
