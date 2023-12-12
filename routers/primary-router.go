package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/squ1ggly/squ1ggly-api-go/controllers"
)

func PrimaryRoutes(primaryRoute string, router *gin.Engine) {
	pRoutes := router.Group(primaryRoute)
	{
		pRoutes.GET("/hello-world", controllers.HelloWorldController)
	}
	TestRoutes(pRoutes)
}
