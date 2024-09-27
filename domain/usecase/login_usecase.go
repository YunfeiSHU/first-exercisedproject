package usecase

import (
	"gin-jwt-gorm/domain/model"
	"gin-jwt-gorm/internal/tokenutil"
)

// loginUsecase结构体 与model层的UserRepository接口,嵌套实现对接
// 利用repository的user的CRUD方法，实现接口，user-->login

type loginUsecase struct {
	userRepository model.UserRepository
}

func NewLoginUsecase(userRepository model.UserRepository) model.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}

func (lu *loginUsecase) GetUserByEmail(email string) (*model.User, error) {
	return lu.userRepository.GetByEmail(email)
}

func (lu *loginUsecase) CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
