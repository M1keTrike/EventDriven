package application

import "github.com/M1keTrike/EventDriven/services/rabbitmq/application"

type SendMessageUseCase struct {
	sm_s *application.ReturnOfferService
}

func NewSendMessageUseCase(sm_s *application.ReturnOfferService) *SendMessageUseCase {
	return &SendMessageUseCase{sm_s: sm_s}
}

func (s *SendMessageUseCase) Execute(queue string, msg []byte) error {

	s.sm_s.Execute(queue, msg)
	return nil
}
