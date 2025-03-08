package models

type MessageWS struct {
	Sender        string
	DestinationID string
	Content       string
	Time          string
}

func NewMessageWS(sender string, destinationID string, content string, time string) *MessageWS {
	return &MessageWS{
		Sender:        sender,
		DestinationID: destinationID,
		Content:       content,
		Time:          time,
	}
}
