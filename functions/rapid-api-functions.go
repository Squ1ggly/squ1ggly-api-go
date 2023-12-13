package functions

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"os"
	"squ1ggly/squ1ggly-api-go/types"
	"strings"
)

func generateBaseURL(siteStub *types.RapidSiteStub) *string {
	//Local
	if siteStub.Environment == "local" || siteStub.Environment == "localhost:8080" {
		localBaseUrl := "http://localhost:8080/" + url.QueryEscape(siteStub.Tenant) + "/" + url.QueryEscape(siteStub.Site)
		return &localBaseUrl
	}
	//Test
	if siteStub.Environment == "test" {
		testBaseUrl := "https://api-test.rapidplatform.com/api/" + url.QueryEscape(siteStub.Tenant) + "/" + url.QueryEscape(siteStub.Site)
		return &testBaseUrl
	}
	//App
	if siteStub.Environment == "prod" {
		productionBaseUrl := "https://api.rapidplatform.com/api/" + url.QueryEscape(siteStub.Tenant) + "/" + url.QueryEscape(siteStub.Site)
		return &productionBaseUrl
	}
	//Fallback to given Environment
	fallback := ""
	return &fallback
}

func generateListURL(siteStub *types.RapidSiteStub, listName *string) *string {
	listUrl := *generateBaseURL(siteStub) + "/lists/" + url.QueryEscape(*listName) + "/All$/items"
	return &listUrl
}

func getJsonFromHttpResponse(httpRes *http.Response) *types.Response {
	var dynamicData types.Response

	resBody, err := io.ReadAll(httpRes.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return &dynamicData
	}

	// Unmarshal JSON into the map
	err = json.Unmarshal(resBody, &dynamicData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err, string(resBody))
		return &dynamicData
	}

	return &dynamicData
}

func SetRapidAuth(context *gin.Context) {
	rapidSiteStub := context.MustGet("RapidSiteStub").(*types.RapidSiteStub)
	keyName := "RapidAuthToken" + rapidSiteStub.Site + rapidSiteStub.Tenant + rapidSiteStub.Environment
	storedKey := os.Getenv(keyName)
	if storedKey == "" {
		authUrl := "https://login.microsoftonline.com/" + url.QueryEscape(context.MustGet("RapidSiteStub").(*types.RapidSiteStub).Tenant) + ".onmicrosoft.com/oauth2/token"

		headers := &map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		}

		authCodesBody := ``

		opts := &types.RequestOptions{
			Method:  "POST",
			Headers: *headers,
			Body:    authCodesBody,
		}

		res, err := Fetch(authUrl, *opts)

		if err != nil {
			context.Next()
		}

		var tokenResp types.TokenResponse

		resBody, _ := io.ReadAll(res.Body)

		err = json.Unmarshal(resBody, &tokenResp)
		err = nil
		if err != nil {
			context.Next()
			return
		}
		err = os.Setenv(keyName, tokenResp.TokenType+" "+tokenResp.AccessToken)
		if err != nil {
			return
		}
		err = os.Setenv(keyName, tokenResp.TokenType+" "+tokenResp.AccessToken)
		if err != nil {
			return
		}
		context.Set("RapidAuthToken", tokenResp.TokenType+" "+tokenResp.AccessToken)
	} else {
		context.Set("RapidAuthToken", storedKey)
	}
	context.Next()
}

func GetRapidItems(context *gin.Context, listName *string, filter *string) *types.Response {
	var dynamicData types.Response

	rapidSiteStub := context.MustGet("RapidSiteStub").(*types.RapidSiteStub)
	token := context.MustGet("RapidAuthToken").(string)

	listUrl := *generateListURL(rapidSiteStub, listName)
	//listUrl += "?$select='*'"
	if *filter != "" {
		if strings.Contains(listUrl, "?") {
			listUrl += "&$filter=" + url.QueryEscape(*filter)
		} else {
			listUrl += "?$filter=" + url.QueryEscape(*filter)
		}
	}

	headers := map[string]string{
		"Authorization": token,
	}

	opts := &types.RequestOptions{
		Method:  "GET",
		Headers: headers,
	}

	response, err := Fetch(listUrl, *opts)
	if err != nil {
		fmt.Println("Error:", err)
		return &dynamicData
	}
	keyName := "RapidAuthToken" + rapidSiteStub.Site + rapidSiteStub.Tenant + rapidSiteStub.Environment
	if response.StatusCode == 400 {
		err := os.Setenv(keyName, "")
		if err != nil {
			return nil
		}
		SetRapidAuth(context)
		response, err = Fetch(listUrl, *opts)
		if err != nil {
			fmt.Println("Error:", err)
			return &dynamicData
		}
	}

	dynamicData = *getJsonFromHttpResponse(response)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	return &dynamicData
}
