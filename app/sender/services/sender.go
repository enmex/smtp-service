package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"smtp/app/config"
	"smtp/app/sender/dto"
	"smtp/pkg/logger"
)

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) Send(payload dto.SendMailPayload) error {
	logger.Logger.Info(config.GetConfig().Mode)

	providerName := payload.Provider
	var provider config.Provider
	var exist bool
	if providerName == nil {
		provider, exist = config.GetConfig().Providers["default"]
	} else {
		provider, exist = config.GetConfig().Providers[*providerName]
	}
	if !exist {
		return fmt.Errorf("%s provider not found", *providerName)
	}

	host := provider.Delivery.Host
	addr := provider.Delivery.Address

	auth := smtp.PlainAuth("", payload.SenderMail, provider.Credentials.Password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(payload.SenderMail); err != nil {
		return err
	}

	if err = c.Rcpt(payload.Recipient); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	message := s.buildMessage(payload, payload.SenderMail)

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}

func (s *Sender) buildMessage(payload dto.SendMailPayload, sender string) string {
	delimiter := "beer=happiness"

	msg := fmt.Sprintf("From: %s\r\n", sender)
	msg += fmt.Sprintf("To: %s\r\n", payload.Recipient)
	msg += fmt.Sprintf("Subject: %s\r\n", payload.Subject)
	msg += fmt.Sprintf("MIME-version: 1.0;\nContent-Type: multipart/mixed; charset=\"UTF-8\"; boundary=%s\r\n", delimiter)
	msg += fmt.Sprintf("\r\n%s\r\n", payload.Message)

	//message
	msg += fmt.Sprintf("\r\n--%s\r\n", delimiter)
	msg += fmt.Sprintf("Content-Type: text/html; charset=\"utf-8\"\r\n")
	msg += fmt.Sprintf("\r\n%s\r\n", payload.Message)

	//attachments
	if payload.Attachments != nil {
		msg += fmt.Sprintf("\r\n--%s\r\n", delimiter)
		for attachmentName, attachment := range *payload.Attachments {
			msg += fmt.Sprintf("\r\n--%s\r\n", delimiter)
			msg += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
			msg += "Content-Transfer-Encoding: base64\r\n"
			msg += "Content-Disposition: attachment;filename=\"" + attachmentName + "\"\r\n"
			msg += "\r\n" + attachment
		}
	}

	return msg
}
