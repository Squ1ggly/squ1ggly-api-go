package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/squ1ggly/squ1ggly-api-go/middleware"
)

func TestController(context *gin.Context) {
	rapidSiteStub := context.MustGet("rapidSiteStub").(*middleware.RapidSiteStub)
	context.IndentedJSON(http.StatusOK, rapidSiteStub)
}
