package middleware

import (
	"encoding/json"
	"io"
	"net/url"
	"squ1ggly/squ1ggly-api-go/functions"
	"squ1ggly/squ1ggly-api-go/types"

	"github.com/gin-gonic/gin"
)

func SetRapidAuth(context *gin.Context) {

	url := "https://login.microsoftonline.com/" + url.QueryEscape(context.MustGet("RapidSiteStub").(*types.RapidSiteStub).Tenant) + ".onmicrosoft.com/oauth2/token"

	headers := &map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	authCodesBody := ""

	opts := &types.RequestOptions{
		Method:  "POST",
		Headers: *headers,
		Body:    authCodesBody,
	}

	res, err := functions.Fetch(url, *opts)

	if err != nil {
		context.Next()
	}

	var tokenResp types.TokenResponse

	resBody, _ := io.ReadAll(res.Body)

	json.Unmarshal(resBody, &tokenResp)

	context.Set("RapidAuthToken", tokenResp.TokenType+" "+tokenResp.AccessToken)
	context.Next()
}
