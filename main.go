package counter

type Name string

type CountService interface {
	// Increases the counter by one.
	Increase(name Name) error

	// Gets the current value for the counter. If no counter
	// exists with the specified id zero is returned.
	Get(name Name) (int, error)
}
