package mailer

import (
	"github.com/gobuffalo/packr"
	"github.com/jinzhu/configor"
	gomail "gopkg.in/gomail.v2"
)

// Mailer mail struct interface
type Mailer interface {
	Send(recipient Recipienter, mailTplPath string, confCode string) error
	config() Configer
	box() packr.Box
	client() *gomail.Dialer
}

type mail struct {
	c  Configer
	cl *gomail.Dialer
	b  packr.Box
}

// Send sends email
func (m *mail) Send(recipient Recipienter, mailTplPath string, confCode string) error {
	mailTpl := &mailTemplate{}
	if err := configor.Load(mailTpl, mailTplPath); err != nil {
		return err
	}
	mailTpl.UserID = recipient.PrimaryKey()
	mailTpl.Email = recipient.EmailAddress()
	mailTpl.Name = recipient.Name()
	mailTpl.code = confCode

	mailBody, err := parseTemplate(m, mailTpl)
	if err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetAddressHeader("From", m.config().SenderEmail(), m.config().SenderName())
	msg.SetAddressHeader("To", recipient.EmailAddress(), recipient.Name())
	msg.SetHeader("Subject", replacePlaceholders(mailTpl.Subject, mailTpl))
	msg.SetBody("text/html", mailBody)

	if err = m.client().DialAndSend(msg); err != nil {
		return err
	}

	return nil
}

// returns mailer config
func (m *mail) config() Configer {
	return m.c
}

// returns mailer client
func (m *mail) client() *gomail.Dialer {
	return m.cl
}

// returns mailer storage box
func (m *mail) box() packr.Box {
	return m.b
}

// Recipienter recipient interface
type Recipienter interface {
	PrimaryKey() string
	Name() string
	EmailAddress() string
}
