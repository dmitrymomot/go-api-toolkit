package application

import (
	"fmt"

	"github.com/dmitrymomot/go-api-toolkit/helpers"
)

// Configer config interface
type Configer interface {
	APIVersion() string
	APIMajorVersion() string
	Env() string
	BaseURL() string
	MasterPassword(password string) bool
	MinAllowedUserAge() int
	EmailOfSupport() string
	ProductName() string
	ProductLink() string
	RefCouponAmount() uint
	RefCouponPrefix() string
}

// Config of the package
type Config struct {
	Version *struct {
		Major int
		Minor int
		Patch int
	}
	AppEnv             string
	APIBaseURL         string
	MasterPasswordHash string
	MinUserAge         int
	SupportEmail       string
	Product            *struct {
		Name string
		Link string
	}
	RefCoupon *struct {
		Prefix string
		Amount uint
	}
}

// APIVersion returns current API version
func (c *Config) APIVersion() string {
	if c.Version == nil {
		return ""
	}
	return fmt.Sprintf("%d.%d.%d", c.Version.Major, c.Version.Minor, c.Version.Patch)
}

// APIMajorVersion returns major number of current API version
func (c *Config) APIMajorVersion() string {
	if c.Version == nil {
		return ""
	}
	return fmt.Sprintf("v%d", c.Version.Major)
}

// Env returns current application environment
func (c *Config) Env() string {
	return c.AppEnv
}

// BaseURL return base app url
func (c *Config) BaseURL() string {
	return c.APIBaseURL
}

// MasterPassword checks whether password is master password
func (c *Config) MasterPassword(password string) bool {
	return helpers.CheckHash(c.MasterPasswordHash, password)
}

// MinAllowedUserAge returns minimal user age
func (c *Config) MinAllowedUserAge() int {
	return c.MinUserAge
}

// EmailOfSupport returns support email address
func (c *Config) EmailOfSupport() string {
	return c.SupportEmail
}

// ProductName returns product name
func (c *Config) ProductName() string {
	if c.Product == nil {
		return ""
	}
	return c.Product.Name
}

// ProductLink returns product link
func (c *Config) ProductLink() string {
	if c.Product == nil {
		return ""
	}
	return c.Product.Link
}

// RefCouponAmount returns referral coupon amount
func (c *Config) RefCouponAmount() uint {
	if c.RefCoupon == nil {
		return 0
	}
	return c.RefCoupon.Amount
}

// RefCouponPrefix returns referral coupon prefix
func (c *Config) RefCouponPrefix() string {
	if c.RefCoupon == nil {
		return ""
	}
	return c.RefCoupon.Prefix
}
