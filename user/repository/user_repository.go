package repository

import (
	"gorm.io/gorm"
	"swn_realtime_vote/domain"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{DB: db}
}

func (u userRepository) AddUser(user *domain.User) (us *domain.User, err error) {
	err = u.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}
