package controllers

import (
	"net/http"
	"squ1ggly/squ1ggly-api-go/types"

	"github.com/gin-gonic/gin"
)

func TestController(context *gin.Context) {
	rapidSiteStub := context.MustGet("RapidSiteStub").(*types.RapidSiteStub)
	context.IndentedJSON(http.StatusOK, rapidSiteStub)
}
