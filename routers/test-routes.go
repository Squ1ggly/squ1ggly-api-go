package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/squ1ggly/squ1ggly-api-go/controllers"
)

func TestRoutes(pRoutes *gin.RouterGroup) {
	tRoute := pRoutes.Group("/test")
	{
		tRoute.GET("/1", controllers.TestController)
	}
}
