package controllers

import (
	usecases "github.com/M1keTrike/EventDriven/offers/application/UseCases"
	"github.com/gin-gonic/gin"
)

type CreateOfferRequest struct {
	UserId  int     `json:"user_id"`
	Product string  `json:"product"`
	Price   float32 `json:"price"`
	Time    string  `json:"time"`
}

type CreateOfferController struct {
	co usecases.CreateOfferUseCase
}

func NewCreateOfferController(co usecases.CreateOfferUseCase) *CreateOfferController {
	return &CreateOfferController{
		co: co,
	}
}

func (co_c *CreateOfferController) Execute(c *gin.Context) {

	var req CreateOfferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := co_c.co.Execute(req.UserId, req.Product, req.Price, req.Time)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Offer created successfully"})
}
