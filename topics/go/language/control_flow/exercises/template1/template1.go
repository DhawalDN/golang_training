// Write a program that inspects a user's name and greets them in a certain way
// if they are on a list or in a different way if they are not. Also look at
// the user's age and tell them some special secret if they are old enough to
// know it.
package main

// Add imports
import (
	"fmt"
)

func main() {

	// Change these values and rerun your program.
	name := "Moe"
	age := 22

	// If the user's name is on a special list then give them a secret greeting.
	switch name {
	case "John", "Bella", "Jack", "Moe":
		fmt.Println("You are a hero")
	}

	// If the user is old enough then tell them a secret.
	switch {
	case age > 18:
		{
			fmt.Println("You are eligible")
		}
	default:
		fmt.Println("Can't tell you any secret")
	}

}
