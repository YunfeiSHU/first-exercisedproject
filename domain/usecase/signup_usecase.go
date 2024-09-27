package usecase

import (
	"gin-jwt-gorm/domain/model"
	"gin-jwt-gorm/internal/tokenutil"
)

type signupUsecase struct {
	userRepository model.UserRepository
}

func NewSignupUsecase(userRepository model.UserRepository) *signupUsecase {
	return &signupUsecase{userRepository}
}

func (su *signupUsecase) Create(user *model.User) error {
	return su.userRepository.Create(user)
}
func (su *signupUsecase) GetUserByEmail(email string) (*model.User, error) {
	return su.userRepository.GetByEmail(email)
}
func (su *signupUsecase) CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}
func (su *signupUsecase) CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
