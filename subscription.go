package contracts

// Subscription describes a set of shows to fetch and the episode limit
// to request for each of them.
type Subscription interface {
	Limit() int
	Shows() []string
}
