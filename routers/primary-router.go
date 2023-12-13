package routers

import (
	"squ1ggly/squ1ggly-api-go/controllers"

	"github.com/gin-gonic/gin"
)

func PrimaryRoutes(primaryRoute string, router *gin.Engine) {
	pRoutes := router.Group(primaryRoute)
	{
		pRoutes.GET("/hello-world", controllers.HelloWorldController)
	}
	RapidRoutes(pRoutes)
	TestRoutes(pRoutes)
}
