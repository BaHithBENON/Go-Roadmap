package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting message for a given name.
//
// Parameters:
// name: A string representing the name of the person to greet.
//
// Returns:
// A string representing the greeting message.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name")
	}

	// Generate the greeting message using the provided name.
	// The message is formatted using the fmt.Sprintf function.
	message := fmt.Sprintf("Hello, %s. Welcome in my app!", name)

	// Return the generated greeting message.
	return message, nil
}
