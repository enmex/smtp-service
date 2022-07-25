package dto

type SendMailPayload struct {
	Provider    string             `json:"provider,omitempty"`
	SenderMail  string             `json:"sender_mail"`
	Subject     string             `json:"subject"`
	Message     string             `json:"message"`
	Recipient   string             `json:"recipient"`
	Attachments *map[string]string `json:"attachments,omitempty"`
}
