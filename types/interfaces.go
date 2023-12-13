package types

import "github.com/gin-gonic/gin"

type CustomContext struct {
	*gin.Context
	RapidSiteStub  string `json:"rapid_site_stub"`
	RapidAuthToken string
}
