package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type APIClient struct {
	client *http.Client
	url    string
}

func NewAPIClient(url string, timeout time.Duration) *APIClient {
	return &APIClient{
		client: &http.Client{
			Timeout: timeout,
		},
		url: url,
	}
}

func (c *APIClient) PostWeather(payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 500 {
		return errors.New("server error, retrying later")
	}

	if res.StatusCode >= 400 {
		return errors.New("bad request: " + res.Status)
	}

	return nil
}
