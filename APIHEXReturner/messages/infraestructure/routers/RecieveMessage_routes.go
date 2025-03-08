package routers

import (
	"github.com/M1keTrike/EventDriven/messages/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func AttachRecieveRoutes(r *gin.Engine, recieveController *controllers.RecieveMessageController) {
	r.POST("/recieve", recieveController.Execute)

}
