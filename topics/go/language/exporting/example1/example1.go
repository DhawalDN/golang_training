// Sample program to show how to access an exported identifier.
package main

import (
	"fmt"

	"golang_training/topics/go/language/exporting/example1/counters"
)

func main() {

	// Create a variable of the exported type and initialize the value to 10.
	counter := counters.AlertCounter(50)

	fmt.Printf("Counter: %d\n", counter)
}
