package termii

const (
	EndpointSendToken   = "api/sms/otp/send"
	EndpointVoiceToken  = "api/sms/otp/send/voice"
	EndpointVoiceCall   = "api/sms/otp/call"
	EndpointInAppToken  = "api/sms/otp/generate"
	EndpointVerifyToken = "api/sms/otp/verify"
	EndpointEmailToken  = "api/email/otp/send"
)

type MessageType string

const (
	Numeric      = MessageType("NUMERIC")
	AlphaNumeric = MessageType("ALPHANUMERIC")
)

type Channel string

const (
	ChannelEmail    = Channel("email")
	ChannelDND      = Channel("dnd")
	ChannelWhatsapp = Channel("WhatsApp")
	ChannelGeneric  = Channel("generic")
)