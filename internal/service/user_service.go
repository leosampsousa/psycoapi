package service

import (
	"context"

	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepo: ur}
}

func (us *UserService) GetUser(ctx context.Context, username string) (*dto.UserDTO, error) {
	user, err := us.userRepo.GetUser(ctx, username)
	if (err != nil) {
		return nil, err
	}

	return &dto.UserDTO{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Username: user.Username,
	}, nil
}

func (us *UserService) CreateUser(dto dto.UserDTO) {
}