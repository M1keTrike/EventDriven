package dependencies

import (
	"os"

	"github.com/M1keTrike/EventDriven/messages/application"
	"github.com/M1keTrike/EventDriven/messages/infraestructure/controllers"
	"github.com/M1keTrike/EventDriven/messages/infraestructure/routers"
	service "github.com/M1keTrike/EventDriven/services/rabbitmq/application"
	"github.com/M1keTrike/EventDriven/services/rabbitmq/infraestructure"

	"github.com/gin-gonic/gin"
)

type MessageDependencies struct{}

func NewMessageDependencies() *MessageDependencies {
	return &MessageDependencies{}
}

func (d *MessageDependencies) Execute(r *gin.Engine) {
	rbbtURL := os.Getenv("RABBITMQ_URL")

	RabbitMQ, err := infraestructure.NewRabbitMQBus(rbbtURL)
	if err != nil {
		panic(err)
	}

	returnOfferService := service.NewReturnOfferService(RabbitMQ)
	sendMessageUseCase := application.NewSendMessageUseCase(returnOfferService)
	recieveMessageuseCase := application.NewRecieveMessageUseCase(sendMessageUseCase)

	recieveController := controllers.NewRecieveMessageController(*recieveMessageuseCase)
	routers.AttachRecieveRoutes(r, recieveController)

}
