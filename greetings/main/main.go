package main

import (
	"fmt"
	"log"

	"kirin.com/greetings" // Import local package
)

func main() {

	// Logging purpose
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Call the Hello function from the greetings package
	message, err := greetings.Hello("Kirin")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
