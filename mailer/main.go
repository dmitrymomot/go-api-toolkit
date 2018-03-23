package mailer

import (
	gomail "gopkg.in/gomail.v2"
)

// Setup mailer
func Setup(config Configer) Mailer {
	client := gomail.NewPlainDialer(
		config.Host(),
		config.Port(),
		config.Username(),
		config.Password(),
	)
	return &mail{config, client}
}
