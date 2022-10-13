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

func (u *userRepository) isUserExists(phoneNumber string) bool {
	var user domain.User
	err := u.DB.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

func (u *userRepository) AddUser(user *domain.User, ipAddress string) (us *domain.User, err error) {
	if u.isUserExists(user.PhoneNumber) {
		return nil, err
	}
	user.IpAddress = ipAddress
	err = u.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetUser(id int) (us *domain.User, err error) {
	user := domain.User{}
	err = u.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (u *userRepository) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	err := u.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
