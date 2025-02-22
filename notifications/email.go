package notifications

import (
	gomail "gopkg.in/mail.v2"
)

// Should be an interface at some point, for
// multiple notification channels
type Mailer struct {
	dialer *gomail.Dialer
}

func NewMailer(smtpHost string) *Mailer {
	// Very barebones for now
	dialer := gomail.NewDialer(smtpHost, 25, "", "")
	return &Mailer{
		dialer,
	}
}

// We should panic if the program can't send error
// messages, but I'll leave that to the main loop.
func (m *Mailer) SendError(err error) error {

}
