package model

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileUsecase interface {
	GetProfileByID(userID int) (*Profile, error)
}
