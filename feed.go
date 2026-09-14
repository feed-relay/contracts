package contracts

// Feed describes feed metadata
type Feed interface {
	// Limit items per show
	Limit() int

	// Shows returns show list
	Shows() []string

	// Link returns feed's external link
	Link() string

	// Slug returns feed code
	Slug() string

	// Title returns feed title
	Title() string

	// Description returns description
	Description() string

	// Image returns image url
	Image() string
}
