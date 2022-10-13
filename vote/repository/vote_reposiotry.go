package repository

import (
	"gorm.io/gorm"
	"swn_realtime_vote/domain"
)

type voteRepository struct {
	DB *gorm.DB
}

func NewVoteRepository(db *gorm.DB) domain.VoteRepository {
	return &voteRepository{DB: db}
}

func (v voteRepository) AddVote(vote *domain.Vote) (*domain.Vote, error) {
	err := v.DB.Create(vote).Error
	if err != nil {
		return nil, err
	}
	return vote, nil
}
