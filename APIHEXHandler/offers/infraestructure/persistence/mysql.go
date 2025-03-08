package persistence

import (
	"database/sql"

	"github.com/M1keTrike/EventDriven/offers/domain"
)

type OfferRepository struct {
	DB *sql.DB
}

func NewOfferRepository(db *sql.DB) *OfferRepository {
	return &OfferRepository{
		DB: db,
	}
}

func (o *OfferRepository) SaveOffer(offer *domain.Offer) error {
	_, err := o.DB.Exec("INSERT INTO offers (user_id, product, price, time) VALUES (?, ?, ?, ?)", offer.UserId, offer.Product, offer.Price, offer.Time)
	if err != nil {
		return err
	}
	return nil
}

func (o *OfferRepository) GetOffers() ([]*domain.Offer, error) {
	rows, err := o.DB.Query("SELECT * FROM offers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	offers := []*domain.Offer{}
	for rows.Next() {
		offer := &domain.Offer{}
		err := rows.Scan(&offer.Id, &offer.UserId, &offer.Product, &offer.Price, &offer.Time)
		if err != nil {
			return nil, err
		}
		offers = append(offers, offer)
	}
	return offers, nil
}
