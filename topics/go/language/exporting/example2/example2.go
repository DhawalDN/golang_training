// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"

	_ "golang_training/topics/go/language/exporting/example2/counters"
)

// Note : remove '_': it is used for ignoring imports
func main() {

	// Create a variable of the unexported type and initialize the value to 10.
	counter := counters.alertCounter(10)

	// ./example2.go:17: cannot refer to unexported name counters.alertCounter
	// ./example2.go:17: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
