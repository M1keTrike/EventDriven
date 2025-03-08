package routers

import (
	"github.com/M1keTrike/EventDriven/offers/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func AttachOfferRoutes(r *gin.Engine, createOfferController *controllers.CreateOfferController, getOffersController *controllers.GetOffersController) {
	r.POST("/offers", createOfferController.Execute)
	r.GET("/offers", getOffersController.Execute)

}
