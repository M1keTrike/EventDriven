package usecases

import (
	"fmt"

	"github.com/M1keTrike/EventDriven/offers/domain"
	"github.com/M1keTrike/EventDriven/offers/infraestructure/persistence"
	"github.com/M1keTrike/EventDriven/services/rabbitmq/application"
)

type CreateOfferUseCase struct {
	db         persistence.OfferRepository
	messageBus application.PublishOfferService
}

func NewCreateOfferUseCase(db persistence.OfferRepository, mb application.PublishOfferService) *CreateOfferUseCase {
	return &CreateOfferUseCase{
		db:         db,
		messageBus: mb,
	}
}

func (c *CreateOfferUseCase) Execute(userId int, product string, price float32, time string) error {
	offer := domain.NewOrder(userId, product, price, time)
	fmt.Println(offer)
	ofJSON, err := offer.ToJSON()
	fmt.Println(ofJSON)
	if err != nil {
		return err
	}
	c.messageBus.Execute("offers_queue", ofJSON)
	err = c.db.SaveOffer(offer)
	if err != nil {
		return err
	}
	return nil
}
