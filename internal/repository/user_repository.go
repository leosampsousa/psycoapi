package repository

import (
	"context"

	"github.com/leosampsousa/psycoapi/internal/db"
)

type UserRepository struct {
	db *db.Queries
}

func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUser(ctx context.Context, username string) (*db.User, error) {
	user, err := ur.db.GetUser(ctx, username)
	if (err != nil) {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) SaveUser(db.SaveUserParams) {
	
}