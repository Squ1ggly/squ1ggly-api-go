package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type RapidSiteStub struct {
	Environment string `json:"environment"`
	Tenant      string `json:"tenant"`
	Site        string `json:"site"`
}

func SetRapidSiteStub(context *gin.Context) {
	env := strings.ToLower(context.GetHeader("x-minilith-environment"))
	if env == "" {
		env = strings.ToLower(context.Param("env"))
	}

	if env == "production" {
		env = "prod"
	}

	rapidSiteStub := &RapidSiteStub{
		Environment: env,
		Tenant:      strings.ToLower(context.GetHeader("x-minilith-tenant")),
		Site:        strings.ToLower(context.GetHeader("x-minilith-site")),
	}

	context.Set("rapidSiteStub", rapidSiteStub)
	context.Next()
}
