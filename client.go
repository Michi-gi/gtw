package gtw

type GTWClient struct {
	Tweets *TweetsApi
}

func NewClient(apiKey string, apiKeySecret string, accessToken string, accessTokenSecret string, bearerToken string) *GTWClient {
	oauth1Client := NewClientOAuth1(apiKey, apiKeySecret, accessToken, accessTokenSecret)
	oauth2Client := NewClientOAuth2(bearerToken)

	twitter := TweetsApi{
		OAuth1Client: oauth1Client,
		OAuth2Client: oauth2Client,
	}

	client := GTWClient{
		Tweets: &twitter,
	}

	return &client
}

func NewClientOAuth2Only(bearerToken string) *GTWClient {
	oauth2Client := NewClientOAuth2(bearerToken)

	twitter := TweetsApi{
		OAuth2Client: oauth2Client,
	}

	client := GTWClient{
		Tweets: &twitter,
	}

	return &client
}
