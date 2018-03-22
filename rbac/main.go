package rbac

import "github.com/mikespook/gorbac"

// Setup RBAC package
func Setup(config Configer) (RBACer, error) {
	r := &rbac{config, gorbac.New()}
	if err := r.setup(); err != nil {
		panic(err)
	}
	return r, nil
}
