package mailer

import (
	"github.com/gobuffalo/packr"
	gomail "gopkg.in/gomail.v2"
)

// Setup mailer
func Setup(config Configer) Mailer {
	box := packr.NewBox("./templates")
	client := gomail.NewPlainDialer(
		config.Host(),
		config.Port(),
		config.Username(),
		config.Password(),
	)
	return &mail{config, client, box}
}
