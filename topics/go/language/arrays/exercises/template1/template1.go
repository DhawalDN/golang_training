// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var persons [5]string
	// Declare an array of 5 strings and pre-populate it with names.
	users := [5]string{"A", "B", "C", "D", "E"}

	// Assign the populated array to the array of zero values.
	persons[1] = users[1]
	for i, user := range users {
		persons[i] = user
	}
	// Iterate over the first array declared.
	for i := range persons {
		// Display the string value and address of each element.
		fmt.Printf("Persons[%d]= '%v' || address = %p\n", i, persons[i], &persons[i])
	}
}
