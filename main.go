package main

import (
	"squ1ggly/squ1ggly-api-go/middleware"
	routers "squ1ggly/squ1ggly-api-go/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.SetRapidSiteStub)
	routers.PrimaryRoutes("/v1", router)
	router.Run("localhost:3000")
}
