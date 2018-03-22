package storage

import "fmt"

// Configer config interface
type Configer interface {
	Endpoint() string
	AccessKey() string
	SecretKey() string
	Bucket() string
	Region() string
	Host() string
	Secure() bool
	BaseDir() string
}

// Config storage package configuration structure
type Config struct {
	StorageBucket    string
	StorageRegion    string
	StorageHost      string
	StorageAccessKey string
	StorageSecretKey string
	StorageSecure    bool
	StorageBaseDir   string
}

// AccessKey returns storage access key
func (c *Config) AccessKey() string {
	return c.StorageAccessKey
}

// SecretKey returns storage secret key
func (c *Config) SecretKey() string {
	return c.StorageSecretKey
}

// Bucket returns storage bucket name
func (c *Config) Bucket() string {
	return c.StorageBucket
}

// Region returns storage region
func (c *Config) Region() string {
	return c.StorageRegion
}

// Host returns storage host
func (c *Config) Host() string {
	return c.StorageHost
}

// Endpoint returns storage endpoint
func (c *Config) Endpoint() string {
	return fmt.Sprintf("%s.%s", c.Region(), c.Host())
}

// BaseDir returns the base dir into storage bucket
func (c *Config) BaseDir() string {
	return c.StorageBaseDir
}

// Secure returns whether storage has secure connection or not
func (c *Config) Secure() bool {
	return c.StorageSecure
}
