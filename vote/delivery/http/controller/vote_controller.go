package controller

import (
	"github.com/labstack/echo"
	"swn_realtime_vote/domain"
)

type VoteController struct {
	VoteUsecase domain.VoteUsecase
}

func NewVoteController(e *echo.Echo, vu domain.VoteUsecase) {
	controller := &VoteController{VoteUsecase: vu}
	vote := e.Group("api/vote")
	vote.POST("", controller.AddVote)
}

func (vc *VoteController) AddVote(c echo.Context) error {
	var vote domain.Vote
	err := c.Bind(&vote)
	if err != nil {
		return err
	}
	_, err = vc.VoteUsecase.AddVote(&vote)
	if err != nil {
		return err
	}
	return c.JSON(200, vote)
}
