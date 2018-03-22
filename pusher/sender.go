package pusher

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

// Senderer sender struct interface
type Senderer interface {
	Send(clientToken string, title, body string, payload map[string]string) error
}

type sender struct {
	app      *firebase.App
	testMode bool
}

// Send sends push notification to individual device
func (s *sender) Send(clientToken string, title, body string, payload map[string]string) error {
	ctx := context.Background()
	client, err := s.app.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.Message{
		Token: clientToken,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: payload,
	}

	if s.testMode {
		_, err = client.SendDryRun(ctx, message)
		if err != nil {
			return err
		}
	} else {
		_, err = client.Send(ctx, message)
		if err != nil {
			return err
		}
	}

	return nil
}
