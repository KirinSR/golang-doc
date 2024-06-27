package trial

import (
	"fmt"

	"kirin.com/greetings" // Import local package
)

func trial() {
	// Call the Hello function from the greetings package
	message := greetings.Hello("Kirin")
	fmt.Println(message)
}
