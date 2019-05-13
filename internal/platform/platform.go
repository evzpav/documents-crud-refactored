package platform

import (
	"net/http"
	"time"
)

type platformClient struct {
	platformURL string
	configURL   string
	httpClient  internal.HttpClient
}

func NewPlatformClient(platformURL, configURL string) *platformClient {
	return NewWithHttpClient(platformURL, configURL, &http.Client{
		Timeout: 60 * time.Second,
	})
}

func NewWithHttpClient(platformURL, configURL string, client internal.HttpClient) *platformClient {
	return &platformClient{
		platformURL: platformURL,
		configURL:   configURL,
		httpClient:  client,
	}
}