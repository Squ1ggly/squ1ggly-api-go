package main

import (
	"squ1ggly/squ1ggly-api-go/middleware"
	routers "squ1ggly/squ1ggly-api-go/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Use(middleware.SetRapidSiteStub)
	routers.PrimaryRoutes("/v1", router)
	router.Run("localhost:3000")
}
