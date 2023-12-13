package routers

import (
	"squ1ggly/squ1ggly-api-go/controllers"
	"squ1ggly/squ1ggly-api-go/middleware"

	"github.com/gin-gonic/gin"
)

func RapidRoutes(pRoutes *gin.RouterGroup) {
	rRoute := pRoutes.Group("/rapid")
	{
		rRoute.Use(middleware.SetRapidAuthMiddleware)
		rRoute.GET("/get-rapid-items/:listName", controllers.GetItem)
	}
}
