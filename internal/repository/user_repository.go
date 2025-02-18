package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/leosampsousa/psycoapi/internal/db"
	error "github.com/leosampsousa/psycoapi/pkg/errors"
)

type UserRepository struct {
	db *db.Queries
}

func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUserByUsername(ctx context.Context, username string) (*db.User, *error.Error) {
	user, err := ur.db.GetUserByUsername(ctx, username)
	if err == nil {
		return &user, nil
	}

	if err == sql.ErrNoRows {
		return nil, error.NewError(404, "Usuario não encontrado")
	}
	return nil, error.NewError(500, "Erro interno")
}

func (ur *UserRepository) GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (*db.User, *error.Error) {
	user, err := ur.db.GetUserByUsernameAndPassword(ctx, db.GetUserByUsernameAndPasswordParams{Username: username, HashedPassword: password})
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

func (ur *UserRepository) GetFriends(ctx context.Context, userId int32) (*[]db.GetFriendsRow, *error.Error) {
	friends, err := ur.db.GetFriends(ctx, userId)
	if (err != nil) {
		log.Println(err)
		return nil, error.NewError(500, "Não foi possível buscar amigos")
	}

	return &friends, nil
}

//refatorar logica para criar confirmação de amizade
func (ur *UserRepository) AddFriend(ctx context.Context, userId int32, friendId int32) (*error.Error) {
	err := ur.db.AddFriend(ctx, db.AddFriendParams{IDUser: userId, IDFriend: friendId})
	if (err != nil) {
		return error.NewError(500, "Não foi possível adicionar amigo")
	}

	errMirrorRelation := ur.db.AddFriend(ctx, db.AddFriendParams{IDUser: friendId, IDFriend: userId})

	if (errMirrorRelation != nil) {
		//remover relação crianda anteriormente
		return error.NewError(500, "Não foi possível adicionar amigo")
	}

	return nil
}