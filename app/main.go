package main

import (
	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"swn_realtime_vote/domain"
	"swn_realtime_vote/user/delivery/http"
	"swn_realtime_vote/user/delivery/http/middleware"
	"swn_realtime_vote/user/repository"
	"swn_realtime_vote/user/usecase"
)

func main() {
	eo := echo.New()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return
	}
	userRepo := repository.NewUserRepository(db)
	useruc := usecase.NewUserUsecase(userRepo)
	um := middleware.UserMiddleware{}
	http.NewUserController(eo, useruc, um)
	eo.Logger.Fatal(eo.Start("0.0.0.0:8080"))
}
