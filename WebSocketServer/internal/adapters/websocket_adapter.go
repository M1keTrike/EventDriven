package adapters

import (
	"fmt"
	"net/http"

	"github.com/M1keTrike/EventDriven/internal/core"
	"github.com/M1keTrike/EventDriven/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketAdapter struct {
	clients map[string]map[*websocket.Conn]bool
	service *core.MessageService
}

func NewWebSocketAdapter(service *core.MessageService) *WebSocketAdapter {
	return &WebSocketAdapter{
		clients: make(map[string]map[*websocket.Conn]bool),
		service: service,
	}
}

func (ws *WebSocketAdapter) HandleWebSocket(c *gin.Context) {
	destinationID := c.Query("destinationID")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error al conectar WebSocket:", err)
		return
	}

	defer func() {
		conn.Close()
		if destinationID != "" {
			delete(ws.clients[destinationID], conn)
			if len(ws.clients[destinationID]) == 0 {
				delete(ws.clients, destinationID)
			}
		}
	}()

	if destinationID != "" {
		if _, exists := ws.clients[destinationID]; !exists {
			ws.clients[destinationID] = make(map[*websocket.Conn]bool)
		}
		ws.clients[destinationID][conn] = true
		fmt.Printf("Cliente suscrito a %s\n", destinationID)

		for {
			var msg models.Message
			if err := conn.ReadJSON(&msg); err != nil {
				fmt.Println("Emisor desconectado")
				return
			}

			fmt.Printf("Mensaje recibido en el servidor: %+v\n", msg)

			if msg.DestinationID == "" {
				fmt.Println("Error: El mensaje recibido no tiene DestinationID")
				continue
			}

			ws.SendMessage(&msg)
		}

	} else {
		fmt.Println("Cliente conectado como emisor de mensajes")
		for {
			var msg models.Message
			if err := conn.ReadJSON(&msg); err != nil {
				fmt.Println("Emisor desconectado")
				return
			}

			fmt.Printf("Mensaje recibido en el servidor: %+v\n", msg)

			if msg.DestinationID == "" {
				fmt.Println("Error: El mensaje recibido no tiene DestinationID")
				continue
			}

			ws.SendMessage(&msg)
		}

	}
}

func (ws *WebSocketAdapter) SendMessage(msg *models.Message) {
	fmt.Printf("Intentando enviar mensaje a: %s\n", msg.DestinationID)

	if subscribers, exists := ws.clients[msg.DestinationID]; exists {
		for conn := range subscribers {
			fmt.Printf("Enviando mensaje a suscriptor en %s\n", msg.DestinationID)

			if err := conn.WriteJSON(msg); err != nil {
				fmt.Printf("Error enviando mensaje a %s: %v\n", msg.DestinationID, err)
				conn.Close()
				delete(subscribers, conn)
			}
		}
	} else {
		fmt.Printf("No hay suscriptores para %s\n", msg.DestinationID)
	}
}
