package repository

import (
	"context"
	"database/sql"

	"github.com/leosampsousa/psycoapi/internal/db"
	errHandler "github.com/leosampsousa/psycoapi/pkg/errors"
)

type UserRepository struct {
	db *db.Queries
}

func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUser(ctx context.Context, username string) (*db.User, error) {
	user, err := ur.db.GetUser(ctx, username)
	if err == nil {
		return &user, nil
	}

	if err == sql.ErrNoRows {
		return nil, errHandler.RecursoNaoEncontrado
	}
	return nil, errHandler.ErroInterno
}

func (ur *UserRepository) SaveUser(ctx context.Context, saveParams db.SaveUserParams) error {
	err := ur.db.SaveUser(ctx, saveParams)
	return err
}