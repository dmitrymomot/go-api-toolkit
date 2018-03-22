package database

import (
	"fmt"
	"time"
)

// Configer is the database config interface
type Configer interface {
	ConnectionString() string
	ConnectionStringWithDriver() (string, string)
	ConnectionMaxLifetime() time.Duration
	MaxOpenConnections() int
	MaxIdleConnections() int
	Logging() bool
}

// Config is database configuration struct
type Config struct {
	Driver          string `default:"mysql"`
	Host            string
	Username        string
	Password        string
	DBName          string
	LogMode         bool
	Options         map[string]interface{}
	ConnMaxLifetime int64 `default:"0"`
	MaxOpenConns    int   `default:"10"`
	MaxIdleConns    int   `default:"3"`
}

// ConnectionString return database connection string
func (c *Config) ConnectionString() string {
	queryString := ""
	if c.Options != nil {
		for k, v := range c.Options {
			if queryString == "" {
				queryString = fmt.Sprintf("%s=%v", k, v)
			} else {
				queryString = fmt.Sprintf("%s&%s=%v", queryString, k, v)
			}
		}
	}
	return fmt.Sprintf("%s:%s@%s/%s?%s", c.Username, c.Password, c.Host, c.DBName, queryString)
}

// ConnectionStringWithDriver returns connection string with database driver name
func (c *Config) ConnectionStringWithDriver() (string, string) {
	return c.Driver, c.ConnectionString()
}

// ConnectionMaxLifetime returns connection lifetime
func (c *Config) ConnectionMaxLifetime() time.Duration {
	if c.ConnMaxLifetime > 0 {
		return time.Duration(c.ConnMaxLifetime) * time.Second
	}
	return 0
}

// MaxOpenConnections returns number of max open connections
func (c *Config) MaxOpenConnections() int {
	return c.MaxOpenConns
}

// MaxIdleConnections returns number of max idle connections
func (c *Config) MaxIdleConnections() int {
	return c.MaxIdleConns
}

// Logging is a log mode flag
func (c *Config) Logging() bool {
	return c.LogMode
}
