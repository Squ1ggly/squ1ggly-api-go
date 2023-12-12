package main

import (
	"github.com/gin-gonic/gin"
	"github.com/squ1ggly/squ1ggly-api-go/middleware"
	routers "github.com/squ1ggly/squ1ggly-api-go/routers"
)

func main() {
	router := gin.Default()
	router.Use(middleware.SetRapidSiteStub)
	routers.PrimaryRoutes("/v1", router)
	router.Run("localhost:3000")
}
