package application

import (
	"github.com/M1keTrike/EventDriven/services/rabbitmq/domain/repositories"
)

type ReturnOfferService struct {
	messageBus repositories.IMessageBus
}

func NewReturnOfferService(messageBus repositories.IMessageBus) *ReturnOfferService {
	return &ReturnOfferService{
		messageBus: messageBus,
	}
}

func (s *ReturnOfferService) Execute(queue string, msg []byte) error {

	return s.messageBus.Return(queue, msg)
}
