package mailer

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/dmitrymomot/sprig"
)

// Paths to default html templates
const (
	defaultTemplate = "default.html"
	mailgunTemplate = "mailgun.html" // in this template unsubscribe link is mailgun placeholder
)

type templates map[string]string

func (t templates) get(key string) string {
	if tlpFile := t[key]; tlpFile == "" {
		return tlpFile
	}
	return ""
}

var defaultTemplates = templates{
	"default": defaultTemplate,
	"mailgun": mailgunTemplate,
}

type mailTemplate struct {
	UserID  string
	Name    string
	Subject string

	Preheader string
	Intro     []string
	Outro     []string
	Button    *buttonTpl
	Footer    []string

	UnsubscribeLink string
	RemoveEmailLink string

	Product struct {
		Name string
		Link string
	}

	Email string
	code  string
}

// ButtonTpl button template element
type buttonTpl struct {
	Link  string
	Title string
}

func parseTemplate(mail Mailer, mailTpl *mailTemplate) (string, error) {
	baseTpl, err := mail.box().MustString(mail.config().BaseTemplate())
	if err != nil {
		return "", err
	}

	if mailTpl.Product.Name == "" {
		mailTpl.Product.Name = mail.config().ProductName()
	}
	if mailTpl.Product.Link == "" {
		mailTpl.Product.Link = mail.config().ProductLink()
	}

	t := template.Must(template.New("").Funcs(sprig.FuncMap()).Parse(baseTpl))
	var buffer bytes.Buffer
	if err = t.Execute(&buffer, mailTpl); err != nil {
		return "", err
	}

	buffStr := replacePlaceholders(buffer.String(), mailTpl)

	return buffStr, nil
}

func replacePlaceholders(s string, mailTpl *mailTemplate) string {
	s = strings.Replace(s, ":name", mailTpl.Name, -1)
	s = strings.Replace(s, ":user_id", mailTpl.UserID, -1)
	s = strings.Replace(s, ":email", mailTpl.Email, -1)
	s = strings.Replace(s, ":confirmation_code", mailTpl.code, -1)
	s = strings.Replace(s, ":product_name", mailTpl.Product.Name, -1)
	s = strings.Replace(s, ":product_link", mailTpl.Product.Link, -1)
	return s
}
