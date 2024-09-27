package usecase

import (
	"gin-jwt-gorm/domain/model"
	"gin-jwt-gorm/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository model.UserRepository
}

func NewRefreshTokenUsecase(userRepository model.UserRepository) model.RefreshTokenUsecase {
	return &refreshTokenUsecase{userRepository: userRepository}
}

func (rtu refreshTokenUsecase) GetUserByID(id int) (*model.User, error) {
	return rtu.userRepository.GetByID(id)
}

func (rtu refreshTokenUsecase) CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu refreshTokenUsecase) CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
