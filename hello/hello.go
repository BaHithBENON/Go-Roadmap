package main

import (
	"fmt"
	"log"

	"examples.com/greetings"
)

// main is the entry point of the program.
//
// It sets the properties of the predefined Logger, including
// the log entry prefix and a flag to disable printing
// the time, source file, and line number.
// Then it calls the Hello function from the greetings package,
// passing an empty string as the name parameter. If an error
// occurs, it logs the error and exits the program. Finally,
// it prints the greeting message returned by the Hello function.
func main() {
	// Set the prefix of the Logger to "greetings: "
	log.SetPrefix("greetings: ")

	// Disable printing of the time, source file, and line number
	log.SetFlags(0)

	// Call the Hello function from the greetings package
	// passing an empty string as the name parameter
	message, err := greetings.Hello("")

	// If an error occurs, log it and exit the program
	if err != nil {
		log.Fatalln(err)
	}

	// Print the greeting message returned by the Hello function
	fmt.Println(message)
}
