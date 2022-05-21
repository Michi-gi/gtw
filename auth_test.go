package gtw

import "testing"

func TestNewClientOAuth1(t *testing.T) {
	client := NewClientOAuth1("apiKey", "apiKeySecret", "accessToken", "accessTokenSecret")

	if client.OAuthVer != "1.0" {
		t.Errorf("ill OAuth Ver: %s", client.OAuthVer)
	}

	if client.Client == nil {
		t.Errorf("client is not set")
	}
}

func TestNewClientOAuth2(t *testing.T) {
	client := NewClientOAuth2("bearerToken")

	if client.OAuthVer != "2.0" {
		t.Errorf("ill OAuth Ver: %s", client.OAuthVer)
	}

	if client.Client == nil {
		t.Errorf("client is not set")
	}

	if client.BearerToken != "bearerToken" {
		t.Errorf("bearer token is not set propery")
	}

}
