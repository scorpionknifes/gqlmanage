package graphql

import (
	"context"

	"github.com/scorpionknifes/gqlmanage/models"
)

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }

func (s *subscriptionResolver) NewEmails(ctx context.Context) (<-chan *models.Email, error) {
	email := make(chan *models.Email, 1)

	return email, nil
}
