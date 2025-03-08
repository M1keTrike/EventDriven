package models

type Message struct {
	Sender        string `json:"Sender"`
	DestinationID string `json:"DestinationID"`
	Content       string `json:"Content"`
	Time          string `json:"Time"`
}
