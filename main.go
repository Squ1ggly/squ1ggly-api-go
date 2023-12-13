package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"squ1ggly/squ1ggly-api-go/middleware"
	routers "squ1ggly/squ1ggly-api-go/routers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Use(middleware.SetRapidSiteStub)
	routers.PrimaryRoutes("/v1", router)
	err := router.Run("localhost:3000")
	if err != nil {
		return
	}
}
