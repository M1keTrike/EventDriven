package application

import (
	"fmt"

	"github.com/M1keTrike/EventDriven/messages/domain"
)

type RecieveMessageUseCase struct {
	sm_uc *SendMessageUseCase
}

func NewRecieveMessageUseCase(sm_uc *SendMessageUseCase) *RecieveMessageUseCase {
	return &RecieveMessageUseCase{sm_uc: sm_uc}
}

func (r *RecieveMessageUseCase) Execute(id int, user_id int, product string, price float32, time string) error {

	message := domain.NewMessage(id, user_id, product, price, time)

	fmt.Println(message)

	ofJSON, err := domain.ToJSON(message)

	if err != nil {
		return err
	}
	r.sm_uc.Execute("return_queue", ofJSON)

	return nil
}
