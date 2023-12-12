package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorldController(context *gin.Context) {
	context.String(http.StatusOK, "Hello world")
}
