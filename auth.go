package gtw

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/oauth1"
)

type AuthClient struct {
	OAuthVer    string
	Client      *http.Client
	BearerToken string
}

func NewClientOAuth1(apiKey string, apiKeySecret string, accessToken string, accessTokenSecret string) *AuthClient {
	config := oauth1.NewConfig(apiKey, apiKeySecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	client := AuthClient{
		OAuthVer: "1.0",
		Client:   config.Client(oauth1.NoContext, token),
	}

	return &client
}

func NewClientOAuth2(bearerToken string) *AuthClient {
	client := AuthClient{
		OAuthVer:    "2.0",
		Client:      &http.Client{},
		BearerToken: bearerToken,
	}

	return &client
}

func (a *AuthClient) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	switch a.OAuthVer {
	case "1.0":
		return http.NewRequest(method, url, body)

	case "2.0":
		req, err := http.NewRequest(method, url, body)
		if req != nil {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.BearerToken))
		}

		return req, err

	default:
		return nil, fmt.Errorf("unknown OAuth ver: %s", a.OAuthVer)
	}
}

func (a *AuthClient) Exec(req *http.Request) (*http.Response, error) {
	return a.Client.Do(req)
}
