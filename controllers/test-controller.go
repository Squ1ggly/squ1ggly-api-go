package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"squ1ggly/squ1ggly-api-go/types"
)

func TestController(context *gin.Context) {
	rapidSiteStub := context.MustGet("RapidSiteStub").(*types.RapidSiteStub)
	context.IndentedJSON(http.StatusOK, rapidSiteStub)
}
