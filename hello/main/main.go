package main

import (
	"fmt"
	"log"

	"kirin.com/greetings" // Import local package
)

func main() {
	// Call the Hello function from the greetings package
	message, err := greetings.Hello("Baddie")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
