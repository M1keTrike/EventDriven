package dependencies

import (
	"database/sql"
	"os"

	usecases "github.com/M1keTrike/EventDriven/offers/application/UseCases"
	"github.com/M1keTrike/EventDriven/offers/infraestructure/controllers"
	"github.com/M1keTrike/EventDriven/offers/infraestructure/persistence"
	"github.com/M1keTrike/EventDriven/offers/infraestructure/routers"
	"github.com/M1keTrike/EventDriven/services/rabbitmq/application"
	"github.com/M1keTrike/EventDriven/services/rabbitmq/infraestructure"
	"github.com/gin-gonic/gin"
)

type OfferDependencies struct {
	DB *sql.DB
}

func NewOfferDependencies(db *sql.DB) *OfferDependencies {
	return &OfferDependencies{
		DB: db,
	}
}

func (d *OfferDependencies) Execute(r *gin.Engine) {
	rbbtURL := os.Getenv("RABBITMQ_URL")

	offerPersistence := persistence.NewOfferRepository(d.DB)

	RabbitMQ, err := infraestructure.NewRabbitMQBus(rbbtURL)
	if err != nil {
		panic(err)
	}

	offerService := application.NewPublishOfferService(RabbitMQ)
	createOfferUseCase := usecases.NewCreateOfferUseCase(*offerPersistence, *offerService)
	gstOffersUseCase := usecases.NewGetOffersUseCase(*offerPersistence)

	createOfferController := controllers.NewCreateOfferController(*createOfferUseCase)
	getOffersController := controllers.NewGetOffersController(*gstOffersUseCase)

	routers.AttachOfferRoutes(r, createOfferController, getOffersController)

}
