package application

import "github.com/M1keTrike/EventDriven/services/rabbitmq/domain/repositories"

type PublishOfferService struct {
	messageBus repositories.IMessageBus
}

func NewPublishOfferService(messageBus repositories.IMessageBus) *PublishOfferService {
	return &PublishOfferService{
		messageBus: messageBus,
	}
}

func (s *PublishOfferService) Execute(queue string, msg []byte) error {
	return s.messageBus.Publish(queue, msg)
}
