package termii

import "net/http"

func (c *Client) SendToken(opts *SendTokenOptions) (*SentToken, error) {
	var st = SentToken{}
	opts.ApiKey = c.config.ApiKey

	_, err := c.request(http.MethodPost, EndpointSendToken, opts, st)
	if err != nil {
		return nil, err
	}

	return &st, nil
}

func (c *Client) VoiceToken(opts *VoiceTokenOptions) (*SentToken, error) {
	var st = SentToken{}
	opts.ApiKey = c.config.ApiKey

	_, err := c.request(http.MethodPost, EndpointVoiceToken, opts, st)
	if err != nil {
		return nil, err
	}

	return &st, nil
}
