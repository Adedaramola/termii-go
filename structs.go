package termii

type request struct {
	ApiKey string `json:"api_key"`
}

type SendTokenOptions struct {
	request
	// Type of message that will be generated and sent as part
	// of the OTP message.
	MessageType MessageType `json:"message_type"`
	// Represents email address or phone number depending on
	// the selected channel.
	To             string  `json:"to"`
	From           string  `json:"from"`
	Channel        Channel `json:"channel"`
	PinAttempts    int     `json:"pin_attempts"`
	PinTTL         int     `json:"pin_time_to_live"`
	PinPlaceholder string  `json:"pin_placeholder"`
	MessageText    string  `json:"message_text"`
}

type SentToken struct {
	PinID     string `json:"pinId"`
	To        string `json:"to"`
	SmsStatus string `json:"smsStatus"`
}

type VoiceTokenOptions struct {
	request
	// Phone number must be in international format
	// E.g 2349043566738
	PhoneNumber string `json:"phone_number"`
	PinAttempts string `json:"pin_attempts"`
	PinTTL      int64  `json:"pin_time_to_live"`
	PintLength  int64  `json:"pin_length"`
}
