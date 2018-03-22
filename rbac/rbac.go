package rbac

import (
	"strings"

	"github.com/mikespook/gorbac"
)

// RBACer is the rbac interface
type RBACer interface {
	IsGranted(role, permission string) bool
	AvailableRoles() []string
	AvailableRolesString(sep string) string
	IsAvailableRole(role string) bool
	IsRoleExist(role string) bool
	setup() error
}

type rbac struct {
	config Configer
	rbac   *gorbac.RBAC
}

// IsGranted checks whether a user is granted this permissions
func (r *rbac) IsGranted(role, permission string) bool {
	return r.rbac.IsGranted(role, gorbac.NewStdPermission(permission), nil)
}

// AvailableRoles returns roles are available to free usage
func (r *rbac) AvailableRoles() []string {
	return r.config.AvailableRoles()
}

// AvailableRolesString returns roles are available to free usage as string with passed separator
func (r *rbac) AvailableRolesString(sep string) string {
	return strings.Join(r.config.AvailableRoles(), sep)
}

// IsAvailableRole checks whether role are available to free usage
func (r *rbac) IsAvailableRole(role string) bool {
	for _, v := range r.config.AvailableRoles() {
		if v == role {
			return true
		}
	}
	return false
}

// IsRoleExist checks whether role exists
func (r *rbac) IsRoleExist(role string) bool {
	if r, _, _ := r.rbac.Get(role); r != nil {
		return true
	}
	return false
}

// setup roles and permissions from loaded config
func (r *rbac) setup() error {
	rp := r.config.RolesPermissions()
	for rl, ps := range rp {
		role := gorbac.NewStdRole(rl)
		for _, p := range ps {
			perm := gorbac.NewStdPermission(p)
			if err := role.Assign(perm); err != nil {
				return err
			}
		}
		r.rbac.Add(role)
	}

	inh := r.config.RolesInheritance()
	for rl, ch := range inh {
		if err := r.rbac.SetParents(rl, ch); err != nil {
			return err
		}
	}

	return nil
}
