package dto

type SendMailPayload struct {
	SenderMail string `json:"sender_mail"`
	SenderPassword string `json:"sender_password"`
	Subject     string             `json:"subject"`
	Message     string             `json:"message"`
	Recipient   string             `json:"recipient"`
	Attachments *map[string]string `json:"attachments,omitempty"`
}
