package mailer

import (
	"io/ioutil"
	"log"
)

// Configer is the mailer config interface
type Configer interface {
	SenderEmail() string
	SenderName() string

	ProductName() string
	ProductLink() string

	Host() string
	Port() int
	Username() string
	Password() string

	BaseTemplate() string
}

// Config email client configuration
type Config struct {
	Sender struct {
		Email string
		Name  string
	}

	Product struct {
		Name string
		Link string
	}

	Connect struct {
		Host     string
		Port     int
		Username string
		Password string
	}

	Template string
}

// SenderName is a getter
func (c *Config) SenderName() string {
	var name string
	if c.Sender.Name != "" {
		name = c.Sender.Name
	}
	if c.Product.Name != "" && name != "" {
		name = name + " from " + c.Product.Name
	}
	if c.Product.Name != "" && name == "" {
		name = c.Product.Name
	}
	return name
}

// SenderEmail is a getter
func (c *Config) SenderEmail() string {
	return c.Sender.Email
}

// ProductName is a getter
func (c *Config) ProductName() string {
	return c.Product.Name
}

// ProductLink is a getter
func (c *Config) ProductLink() string {
	return c.Product.Link
}

// Host is a getter
func (c *Config) Host() string {
	return c.Connect.Host
}

// Port is a getter
func (c *Config) Port() int {
	return c.Connect.Port
}

// Username is a getter
func (c *Config) Username() string {
	return c.Connect.Username
}

// Password is a getter
func (c *Config) Password() string {
	return c.Connect.Password
}

// BaseTemplate is just a getter
func (c *Config) BaseTemplate() string {
	if c.Template == "" {
		return getDefaultTemplateHTML()
	}
	str, err := ioutil.ReadFile(c.Template)
	if err != nil {
		log.Println(err)
		return getDefaultTemplateHTML()
	}
	return string(str)
}
