package controllers

import (
	"net/http"

	"squ1ggly/squ1ggly-api-go/middleware"

	"github.com/gin-gonic/gin"
)

func TestController(context *gin.Context) {
	rapidSiteStub := context.MustGet("rapidSiteStub").(*middleware.RapidSiteStub)
	context.IndentedJSON(http.StatusOK, rapidSiteStub)
}
