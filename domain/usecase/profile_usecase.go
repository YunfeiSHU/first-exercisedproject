package usecase

import "gin-jwt-gorm/domain/model"

type profileUsecase struct {
	userRepository model.UserRepository
}

func NewProfileUsecase(userRepository model.UserRepository) model.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
	}
}
func (pu *profileUsecase) GetProfileByID(userID int) (*model.Profile, error) {
	user, err := pu.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return &model.Profile{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
