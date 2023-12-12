package routers

import (
	"squ1ggly/squ1ggly-api-go/controllers"

	"github.com/gin-gonic/gin"
)

func TestRoutes(pRoutes *gin.RouterGroup) {
	tRoute := pRoutes.Group("/test")
	{
		tRoute.GET("/1", controllers.TestController)
	}
}
