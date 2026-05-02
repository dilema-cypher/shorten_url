package gateways

import (
	"errors"
	"net/http"
	"time"
)

type AuthMeClient struct {
	client  *http.Client
	baseURL string
}

func NewAuthMeClient(baseURL string) *AuthMeClient {
	return &AuthMeClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				IdleConnTimeout:     30 * time.Second,
				DisableCompression: false,
			},
		},
		baseURL: baseURL,
	}
}

func (c *AuthMeClient) DoRequest(token string) error {
	req, err := http.NewRequest(http.MethodPost, c.baseURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}


	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid Token")
	}

	return nil
}


