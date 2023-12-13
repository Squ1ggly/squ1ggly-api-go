package middleware

import (
	"github.com/gin-gonic/gin"
	"squ1ggly/squ1ggly-api-go/functions"
)

func SetRapidAuthMiddleware(context *gin.Context) {
	functions.SetRapidAuth(context)
	context.Next()
}
