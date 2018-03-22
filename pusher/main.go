package pusher

import (
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

// Setup pusher package
func Setup(credFilePath string, testMode bool) (Senderer, error) {
	opt := option.WithCredentialsFile(credFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return &sender{app, testMode}, nil
}
