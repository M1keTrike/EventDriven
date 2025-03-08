package usecases

import (
	"github.com/M1keTrike/EventDriven/offers/domain"
	"github.com/M1keTrike/EventDriven/offers/infraestructure/persistence"
)

type GetOffersUseCase struct {
	db persistence.OfferRepository
}

func NewGetOffersUseCase(db persistence.OfferRepository) *GetOffersUseCase {
	return &GetOffersUseCase{
		db: db,
	}
}

func (g *GetOffersUseCase) Execute() ([]*domain.Offer, error) {
	offers, err := g.db.GetOffers()
	if err != nil {
		return nil, err
	}
	return offers, nil
}
