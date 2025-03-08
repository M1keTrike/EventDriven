package domain

import "encoding/json"

type Offer struct {
	Id      int
	UserId  int
	Product string
	Price   float32
	Time    string
}

func NewOrder(userId int, product string, price float32, time string) *Offer {
	return &Offer{
		UserId:  userId,
		Product: product,
		Price:   price,
		Time:    time,
	}
}

func (o *Offer) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}
