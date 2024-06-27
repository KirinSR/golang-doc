package main

import (
	"fmt"

	"kirin.com/greetings" // Import local package
)

func main() {
	// Call the Hello function from the greetings package
	message := greetings.Hello("Baddie")
	fmt.Println(message)
}
