package rbac

// Configer package config interface
type Configer interface {
	RolesPermissions() map[string][]string
	RolesInheritance() map[string][]string
	AvailableRoles() []string
}

// Config of this package
type Config struct {
	Roles       map[string][]string
	Inheritance map[string][]string
	Available   []string
}

// RolesPermissions returns map of roles with array of permissions
func (c *Config) RolesPermissions() map[string][]string {
	return c.Roles
}

// RolesInheritance returns roles inheritance
func (c *Config) RolesInheritance() map[string][]string {
	return c.Inheritance
}

// AvailableRoles returns roles are available to free usage
func (c *Config) AvailableRoles() []string {
	return c.Available
}
