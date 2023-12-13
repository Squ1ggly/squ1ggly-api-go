package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"squ1ggly/squ1ggly-api-go/types"

	"github.com/gin-gonic/gin"
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
		fmt.Println("Error unmarshaling JSON:", err, string(resBody))
		return &dynamicData
	}

	return &dynamicData
}

func GetRapidItems(context *gin.Context, listName *string) *types.Response {
	var dynamicData types.Response

	rapidSiteStub := context.MustGet("RapidSiteStub").(*types.RapidSiteStub)
	token := context.MustGet("RapidAuthToken").(string)

	url := *generateListURL(rapidSiteStub, listName)

	headers := map[string]string{
		"Authorization": token,
	}

	opts := &types.RequestOptions{
		Method:  "GET",
		Headers: headers,
	}

	response, err := Fetch(url, *opts)
	if err != nil {
		fmt.Println("Error:", err)
		return &dynamicData
	}

	dynamicData = *getJsonFromHttpResponse(response)
	defer response.Body.Close()

	return &dynamicData
}
