package model

type User struct {
	ID       uint   `gorm:"primary_key,AUTO_INCREMENT"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `gorm:"type:varchar(20);unique_index"`
}

type UserRepository interface {
	Create(user *User) error
	GetByName(name string) (*[]User, error)
	GetByEmail(email string) (*User, error)
	GetByID(id int) (*User, error)
}
