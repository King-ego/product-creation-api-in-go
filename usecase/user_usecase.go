package usecase

import (
	"api/model"
	"api/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: repo,
	}
}

func (pu *UserUseCase) GetUsers() ([]model.User, error) {
	return pu.repository.GetUsers()
}
