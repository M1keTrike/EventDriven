package controllers

import (
	"github.com/M1keTrike/EventDriven/messages/application"
	"github.com/gin-gonic/gin"
)

type RecieveMessageController struct {
	rm_uc application.RecieveMessageUseCase
}

type RecieveMessageRequest struct {
	Id      int     `json:"Id"`
	UserId  int     `json:"UserId"`
	Product string  `json:"Product"`
	Price   float32 `json:"Price"`
	Time    string  `json:"Time"`
}

func NewRecieveMessageController(rm_uc application.RecieveMessageUseCase) *RecieveMessageController {
	return &RecieveMessageController{rm_uc: rm_uc}
}

func (r *RecieveMessageController) Execute(c *gin.Context) {
	var req RecieveMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := r.rm_uc.Execute(req.Id, req.UserId, req.Product, req.Price, req.Time)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

}
