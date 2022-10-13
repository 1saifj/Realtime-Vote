package domain

type User struct {
	ID          int    `json:"id"           gorm:"id"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
	IpAddress   string `json:"ip_address"   gorm:"ip_address"`
	MacAddress  string `json:"mac_address"  gorm:"mac_address"`
}

type UserUsecase interface {
	AddUser(user *User) (*User, error)
}

type UserRepository interface {
	AddUser(user *User) (*User, error)
}
