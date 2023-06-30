package termii

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultBaseURL = "https://api.ng.termii.com"

type Config struct {
	ApiKey     string
	BaseURL    string
	HttpClient *http.Client
}

type Client struct {
	client  *http.Client
	config  *Config
	context context.Context
}

func NewClient(ctx context.Context, config *Config) (*Client, error) {
	var client = new(Client)

	if isStringEmpty(config.ApiKey) {
		return client, errors.New("provide an api key")
	}

	if isStringEmpty(config.BaseURL) {
		client.config.BaseURL = defaultBaseURL
	}

	if config.HttpClient == nil {
		client.config.HttpClient = &http.Client{
			Timeout: time.Second * 5,
		}
	}

	return client, nil
}

type Response struct {
	*http.Response
}

func (c *Client) request(method, uri string, payload, v interface{}) (*Response, error) {
	var body = new(bytes.Buffer)

	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}
	u := fmt.Sprintf("%s/%s", c.config.BaseURL, uri)
	if method == http.MethodGet {
		u = fmt.Sprintf("%s?api_key=%s", u, c.config.ApiKey)
	}

	req, err := http.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req = req.WithContext(c.context)

	res, err := c.client.Do(req)
	if err != nil {
		// check if context has been cancelled
		select {
		case <-c.context.Done():
			return nil, c.context.Err()
		default:
		}

		if ue, ok := err.(*url.Error); ok {
			return nil, ue
		}

		// return error as is
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body - %v", err)
	}

	return &Response{res}, nil
}

func isStringEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
