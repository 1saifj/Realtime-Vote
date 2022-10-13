package usecase

import "swn_realtime_vote/domain"

type voteUsecase struct {
	voteRepo domain.VoteRepository
}

func NewVoteUsecase(voteRepo domain.VoteRepository) domain.VoteUsecase {
	return &voteUsecase{voteRepo: voteRepo}
}

func (v voteUsecase) AddVote(vote *domain.Vote) (*domain.Vote, error) {
	res, err := v.voteRepo.AddVote(vote)
	if err != nil {
		return nil, err
	}
	return res, nil
}
