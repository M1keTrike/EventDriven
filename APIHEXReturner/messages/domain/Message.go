package domain

import "encoding/json"

type Message struct {
	Id      int     `json:"Id"`
	UserId  int     `json:"UserId"`
	Product string  `json:"Product"`
	Price   float32 `json:"Price"`
	Time    string  `json:"Time"`
}

func ToJSON(m Message) ([]byte, error) {
	return json.Marshal(m)
}

func NewMessage(id int, userId int, product string, price float32, time string) Message {
	return Message{
		Id:      id,
		UserId:  userId,
		Product: product,
		Price:   price,
		Time:    time,
	}
}
