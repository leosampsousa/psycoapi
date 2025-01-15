package repository

import (
	"context"
	"database/sql"

	"github.com/leosampsousa/psycoapi/internal/db"
	error "github.com/leosampsousa/psycoapi/pkg/errors"
)

type UserRepository struct {
	db *db.Queries
}

func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUser(ctx context.Context, username string) (*db.User, *error.Error) {
	user, err := ur.db.GetUser(ctx, username)
	if err == nil {
		return &user, nil
	}

	if err == sql.ErrNoRows {
		return nil, error.NewError(404, "Usuario não encontrado")
	}
	return nil, error.NewError(500, "Erro interno")
}

func (ur *UserRepository) SaveUser(ctx context.Context, saveParams db.SaveUserParams) *error.Error {
	err := ur.db.SaveUser(ctx, saveParams)
	if (err != nil) {
		return error.NewError(500, "Erro ao salvar usuário")
	}
	return nil
}