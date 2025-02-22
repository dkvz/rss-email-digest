package notifications

import (
	gomail "gopkg.in/mail.v2"
)

// Should be an interface at some point, for
// multiple notification channels
type Mailer struct {
	dialer   *gomail.Dialer
	fromAddr string
	toAddr   string
}

func NewMailer(smtpHost string, fromAddr string, toAddr string) *Mailer {
	// Very barebones for now
	dialer := gomail.NewDialer(smtpHost, 25, "", "")
	return &Mailer{
		dialer,
		fromAddr,
		toAddr,
	}
}

func (m *Mailer) sendEmail(subject string, content string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.fromAddr)
	message.SetHeader("To", m.toAddr)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", content)
	return m.dialer.DialAndSend(message)
}

// We should panic if the program can't send error
// messages, but I'll leave that to the main loop.
func (m *Mailer) SendError(err error) error {
	return m.sendEmail("RSS Feed Notifications - Error", err.Error())
}
