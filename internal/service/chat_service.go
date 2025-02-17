package service

import (
	"context"
	"net/http"

	"github.com/leosampsousa/psycoapi/internal/db"
	"github.com/leosampsousa/psycoapi/internal/repository"
	error "github.com/leosampsousa/psycoapi/pkg/errors"
)

type ChatService struct {
	us *UserService
	cr *repository.ChatRepository
}

func NewChatService(cr *repository.ChatRepository, us *UserService) *ChatService {
	return &ChatService{
		cr: cr,
		us: us,
	}
}

func (cs *ChatService) GetAllChats(ctx context.Context, username string) (*[]db.GetAllChatsRow, *error.Error) {
	user, err := cs.us.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	return cs.cr.GetAllChats(ctx, user.ID)
}

func (cs *ChatService) GetChatMessage(ctx context.Context, userId int32, chatId int32) (*[]db.GetChatMessagesRow, *error.Error) {
	participants, err := cs.cr.GetChatParticipants(ctx, chatId)

	if err != nil {
		return nil, err
	}

	var isParticipant = false
	for _, participant := range *participants {
		if (participant == userId) {
			isParticipant = true
			break
		}
	} 

	if(!isParticipant) {
		return nil, error.NewError(http.StatusForbidden, "Sem permissão para acessar o recurso")
	}

	messages, errMessages := cs.cr.GetChatMessages(ctx, chatId)
	if (errMessages != nil) {
		return nil, errMessages
	}

	return messages, nil
}