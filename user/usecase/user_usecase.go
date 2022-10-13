package usecase

import "swn_realtime_vote/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u userUsecase) AddUser(user *domain.User) (*domain.User, error) {
	res, err := u.userRepo.AddUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
