package core

import (
	"time"

	"github.com/M1keTrike/EventDriven/internal/models"
	"github.com/M1keTrike/EventDriven/internal/ports"
)

type MessageService struct {
	repo ports.MessageRepositoryPort
}

func NewMessageService(repo ports.MessageRepositoryPort) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) ProcessMessage(msg *models.Message) (*models.Message, error) {
	msg.Time = time.Now().Format("2006-01-02 15:04:05")
	err := s.repo.SaveMessage(msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
