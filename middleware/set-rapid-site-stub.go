package middleware

import (
	"squ1ggly/squ1ggly-api-go/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetRapidSiteStub(context *gin.Context) {
	env := strings.ToLower(context.GetHeader("x-minilith-environment"))
	if env == "" {
		env = strings.ToLower(context.Param("env"))
	}

	if env == "production" {
		env = "prod"
	}

	tenant := strings.ToLower(context.GetHeader("x-minilith-tenant"))

	if tenant == "" {
		tenant = strings.ToLower(context.Param("tenant"))
	}

	site := strings.ToLower(context.GetHeader("x-minilith-site"))

	if site == "" {
		site = strings.ToLower(context.Param("site"))
	}

	rapidSiteStub := &types.RapidSiteStub{
		Environment: env,
		Tenant:      tenant,
		Site:        site,
	}

	context.Set("RapidSiteStub", rapidSiteStub)
	context.Next()
}
