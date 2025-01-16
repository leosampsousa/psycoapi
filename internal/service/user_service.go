package service

import (
	"context"

	"github.com/leosampsousa/psycoapi/internal/db"
	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/repository"
	error "github.com/leosampsousa/psycoapi/pkg/errors"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepo: ur}
}

func (us *UserService) GetUser(ctx context.Context, username string, password string) (*dto.UserDTO, *error.Error) {
	user, err := us.userRepo.GetUserByUsernameAndPassword(ctx, username, password)
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

func (us *UserService) CreateUser(ctx context.Context, dto dto.CreateUserDTO) *error.Error {

	if (us.alreadyRegistered(ctx, dto)) {
		return error.NewError(400, "Usuario ja cadastrado")
	}

	err := us.userRepo.SaveUser(
		ctx, 
		db.SaveUserParams{
			FirstName: dto.FirstName, 
			LastName: dto.LastName, 
			Username: dto.Username, 
			HashedPassword: dto.Password, 
		},
	)
	
	return err
}

func (us *UserService) alreadyRegistered(ctx context.Context, dto dto.CreateUserDTO) bool {
	user, _ := us.userRepo.GetUserByUsername(ctx, dto.Username)
	return user != nil 
}