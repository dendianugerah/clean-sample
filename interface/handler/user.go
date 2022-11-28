package handler

import (
	"ddd-sample/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	CreateUser() echo.HandlerFunc
	GetUser() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type reqUser struct {
	Username string `json:"username"`
}

type responseUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (uh *userHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req reqUser
		err := c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdUser, err := uh.userUsecase.Create(req.Username)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		result := responseUser{
			ID:       createdUser.ID,
			Username: createdUser.Username,
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundUser, err := uh.userUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseUser{
			ID:       foundUser.ID,
			Username: foundUser.Username,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundUsers, err := uh.userUsecase.FindAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var res []responseUser

		for _, user := range *foundUsers {
			users := user
			new_user := responseUser{
				ID:       users.ID,
				Username: users.Username,
			}
			res = append(res, new_user)
		}

		return c.JSON(http.StatusOK, res)
	}
}
