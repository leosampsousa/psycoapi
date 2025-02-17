package repository

import (
	"context"

	"github.com/leosampsousa/psycoapi/internal/db"
	error "github.com/leosampsousa/psycoapi/pkg/errors"
)

type ChatRepository struct {
	db *db.Queries
}


func NewChatRepository(db *db.Queries) *ChatRepository{
	return &ChatRepository{db: db}
}

func (cr *ChatRepository) GetAllChats(ctx context.Context, userId int32) (*[]db.GetAllChatsRow, *error.Error) {
	chats, err := cr.db.GetAllChats(ctx, userId)
		
	if (err == nil) {
		return &chats, nil
	}
	
	return nil, error.NewError(500, "Erro interno")	
}

func (cr *ChatRepository) GetChatMessages(ctx context.Context, chatId int32) (*[]db.GetChatMessagesRow, *error.Error) {
	messages, err := cr.db.GetChatMessages(ctx, chatId)
	if err == nil {
		return &messages, nil
	}

	return nil, error.NewError(500, "Erro interno")	
}

func (cr *ChatRepository) GetChatParticipants(ctx context.Context, chatId int32)(*[]int32, *error.Error) {
	participantsIds, err := cr.db.GetChatParticipants(ctx, chatId)

	if (err == nil) {
		return &participantsIds, nil
	}

	return nil, error.NewError(500, "Erro interno")	
}