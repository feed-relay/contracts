package contracts

import (
	"context"

	"github.com/feed-relay/rsscast"
)

type Provider interface {
	// Platform code
	Platform() string

	// Feeds fetches feeds from subscriptions
	Feeds(ctx context.Context, subscriptions []Subscription) (map[string]*rsscast.Feed, error)
}
