package types

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RequestOptions struct {
	Headers map[string]string
	Method  string
	Body    string
}

type RapidSiteStub struct {
	Environment string `json:"environment"`
	Tenant      string `json:"tenant"`
	Site        string `json:"site"`
}

type Response struct {
	Value []interface{} `json:"value"`
}
