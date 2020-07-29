package graphql

import (
	"context"
	"encoding/json"

	"github.com/scorpionknifes/gqlmanage/middleware"
	"github.com/scorpionknifes/gqlmanage/models"
)

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }

func (s *subscriptionResolver) NewEmails(ctx context.Context) (<-chan *models.Email, error) {
	_, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errUnauthenticated
	}
	channel := make(chan *models.Email, 1)
	go func() {
		sub := s.Redis.Subscribe(ctx, "email")
		_, err := sub.Receive(ctx)
		if err != nil {
			return
		}
		ch := sub.Channel()
		for {
			select {
			case message := <-ch:
				var email models.Email
				err := json.Unmarshal([]byte(message.Payload), &email)
				if err != nil {
					sub.Close()
					return
				}
				channel <- &email
			// close when context done
			case <-ctx.Done():
				sub.Close()
				return
			}
		}
	}()
	return channel, nil
}
