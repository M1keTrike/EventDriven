package controllers

import (
	usecases "github.com/M1keTrike/EventDriven/offers/application/UseCases"
	"github.com/gin-gonic/gin"
)

type GetOffersController struct {
	gao usecases.GetOffersUseCase
}

func NewGetOffersController(gao usecases.GetOffersUseCase) *GetOffersController {
	return &GetOffersController{
		gao: gao,
	}
}

func (gao_c *GetOffersController) Execute(c *gin.Context) {
	res, err := gao_c.gao.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"offers": res})
}
