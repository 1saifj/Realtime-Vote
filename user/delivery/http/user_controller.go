package http

import (
	"github.com/labstack/echo"
	"swn_realtime_vote/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(e *echo.Echo, u domain.UserUsecase) {
	controller := &UserController{
		UserUsecase: u,
	}
	e.POST("/user", controller.AddUser)
	//e.GET("/user", validate, uc.GetUser)
}

func (uc *UserController) AddUser(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	err = uc.UserUsecase.AddUser(&user)
	if err != nil {
		return c.JSON(500, ResponseError{Message: err.Error()})
	}
	return c.JSON(200, user)
}
