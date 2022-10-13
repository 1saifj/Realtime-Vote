package domain

type Vote struct {
	ID     int `json:"id"           gorm:"id"`
	UserId int `json:"user_id"      gorm:"user_id"`
}

type VoteUsecase interface {
	AddVote(*Vote) (*Vote, error)
}

type VoteRepository interface {
	AddVote(*Vote) (*Vote, error)
}
