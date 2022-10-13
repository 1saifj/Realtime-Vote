package http

import (
	"github.com/labstack/echo"
	"strconv"
	"swn_realtime_vote/domain"
	"swn_realtime_vote/user/delivery/http/middleware"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserController struct {
	UserUsecase domain.UserUsecase
	UM          middleware.UserMiddleware
}

func NewUserController(e *echo.Echo, u domain.UserUsecase, um middleware.UserMiddleware) {
	controller := &UserController{
		UserUsecase: u,
		UM:          um,
	}

	user := e.Group("api/user")
	user.POST("", controller.AddUser, controller.UM.GetUserData)
	user.GET("/:id", controller.GetUser)
	user.GET("/all", controller.GetAllUsers)
}

func (uc *UserController) AddUser(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	_, err = uc.UserUsecase.AddUser(&user, c.Get("ip").(string))
	if err != nil {
		return err
	}
	if err != nil {
		return c.JSON(500, ResponseError{Message: err.Error()})
	}
	return c.JSON(200, user)
}

func (uc *UserController) GetUser(c echo.Context) error {
	uid := c.Param("id")
	id, err := strconv.Atoi(uid)
	user, err := uc.UserUsecase.GetUser(id)
	if err != nil {
		return c.JSON(500, ResponseError{Message: err.Error()})
	}
	return c.JSON(200, user)

}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.UserUsecase.GetAllUsers()
	if err != nil {
		return c.JSON(500, ResponseError{Message: err.Error()})
	}
	return c.JSON(200, users)
}
