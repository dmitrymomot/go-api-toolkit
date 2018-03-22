package auth

// Setup init jwt
func Setup(c Configer) (JWT, error) {
	return &jsonWebToken{c}, nil
}
