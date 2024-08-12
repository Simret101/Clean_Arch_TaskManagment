package usecase

import (
	"task/domain"
)

type UserUsecase struct {
	UserRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: repo}
}

func (uc *UserUsecase) Register(user *domain.User) error {

	return uc.UserRepo.CreateUser(user)
}

func (uc *UserUsecase) Login(credentials domain.Credentials) (string, error) {

	return uc.UserRepo.AuthenticateUser(credentials.Username, credentials.Password)
}

func (uc *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {

	return uc.UserRepo.GetUserByUsername(username)
}

func (uc *UserUsecase) ValidateToken(tokenString string) (*domain.Claims, error) {

	return uc.UserRepo.ValidateToken(tokenString)
}
