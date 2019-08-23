// Package toy contains support for managing toy inventory.
package toy

// Toy is exported type
type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

// Declare a function named New that accepts values for the
// exported fields. Return a pointer of type Toy that is initialized
// with the parameters.

// Declare a method named OnHand with a pointer receiver that
// returns the current on hand count.

// Declare a method named UpdateOnHand with a pointer receiver that
// updates and returns the current on hand count.

// Declare a method named Sold with a pointer receiver that
// returns the current sold count.

// Declare a method named UpdateSold with a pointer receiver that
// updates and returns the current sold count.
