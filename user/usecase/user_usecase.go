package usecase

import "swn_realtime_vote/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) AddUser(user *domain.User, ipAddress string) (*domain.User, error) {
	//before add check if user already exists
	res, err := u.userRepo.AddUser(user, ipAddress)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userUsecase) GetUser(id int) (*domain.User, error) {
	res, err := u.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userUsecase) GetAllUsers() ([]*domain.User, error) {
	res, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return res, nil
}
