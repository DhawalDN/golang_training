// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import "fmt"

const (
	// Declare a constant named server of kind string and assign a value.
	server = "10.1.10.138"

	// Declare a constant named port of type integer and assign a value.
	port = 8080
)

func main() {

	// Display the value of both server and port.
	fmt.Printf("Server\t: %v \nPort\t: %v\n", server, port)
	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	result := (10 / 1.11111)
	// Display the value of the variable.
	fmt.Println(result)
}
