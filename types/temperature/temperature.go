package temperature

// Temperature represents a floating-point value that can be created as a pointer.
type Temperature *float64

// New creates a pointer to a given temp value.
func New(temp float64) Temperature {
	return Temperature(&temp)
}
