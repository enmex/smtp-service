package sender

type SendMailPayload struct {
	Provider    string            
	SenderMail  string             
	Subject     string             
	Message     string             
	Recipient   string             
	Attachments *map[string]string 
}
