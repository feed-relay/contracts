package contracts

import (
	"context"

	"github.com/feed-relay/rsscast"
)

type Provider interface {
	// Platform code
	Platform() string

	Feeds(ctx context.Context, subscriptions []Subscription) (map[string]*rsscast.Feed, error)
}
