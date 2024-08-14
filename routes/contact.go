package routes

import (
	"github.com/gin-gonic/gin"
	Contact "github.com/singiankay/tsa/controllers"
)

func ContactRoute(router *gin.Engine) {
	contacts := router.Group("/contacts")
	{
		contacts.POST("/", Contact.CreateContact)
		contacts.GET("/", Contact.GetContacts)
		contacts.GET("/:id", Contact.GetContactById)
		contacts.PUT("/:id", Contact.UpdateContact)
		contacts.DELETE("/:id", Contact.DeleteContact)
	}
}