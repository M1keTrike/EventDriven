package sender

import (
	"github.com/gorilla/websocket"
)

type Sender struct {
	url string
}

func NewSender(url string) *Sender {
	return &Sender{
		url: url,
	}
}

func (s *Sender) Send(body []byte, wsconn *websocket.Conn) error {
	err := wsconn.WriteMessage(websocket.TextMessage, body)
	if err != nil {
		return err

	}
	return nil
}
