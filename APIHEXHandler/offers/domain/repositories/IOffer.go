package repositories

import (
	"github.com/M1keTrike/EventDriven/offers/domain"
)

type IOffer interface {
	SaveOffer(offer *domain.Offer) error
	GetOffers() ([]*domain.Offer, error)
}
