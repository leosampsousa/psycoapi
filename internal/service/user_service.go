package service

import (
	"context"

	"github.com/leosampsousa/psycoapi/internal/db"
	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/repository"
	error "github.com/leosampsousa/psycoapi/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepo: ur}
}

func (us *UserService) GetUser(ctx context.Context, username string) (*dto.UserDTO, *error.Error) {
	user, err := us.userRepo.GetUserByUsername(ctx, username)
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

func (us *UserService) Login(ctx context.Context, username string, password string) (*dto.UserDTO, *error.Error) {
	user, err := us.userRepo.GetUserByUsername(ctx, username)
	if (err != nil) {
		return nil, err
	}

	if (!us.verifyPassword(password, user.HashedPassword)) {
		return nil, error.NewError(500, "Usuário ou senha inválidas")
	}

	return us.GetUser(ctx, username)
}

func (us *UserService) CreateUser(ctx context.Context, dto dto.RegisterUserDTO) *error.Error {
	if (us.AlreadyRegistered(ctx, dto.Username)) {
		return error.NewError(400, "Usuario ja cadastrado")
	}

	hashPassword, errHash := us.hashPassword(dto.Password)
	if (errHash != nil) {
		return error.NewError(errHash.Code, errHash.Message)
	}

	err := us.userRepo.SaveUser(
		ctx, 
		db.SaveUserParams{
			FirstName: dto.FirstName, 
			LastName: dto.LastName, 
			Username: dto.Username, 
			HashedPassword: hashPassword, 
		},
	)
	
	return err
}

func (us *UserService) AlreadyRegistered(ctx context.Context, username string) bool {
	user, _ := us.userRepo.GetUserByUsername(ctx, username)
	return user != nil 
}

func (us *UserService) hashPassword(password string) (string, *error.Error) {
	costFactor := 14
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), costFactor)
	if (err != nil) {
		return "", error.NewError(500, "não foi possível criar uma conta")
	}
    return string(bytes), nil
}

func (us *UserService) verifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}