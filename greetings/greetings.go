package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message
	

	// if name is empty
	if name == "" {
		return "", errors.New("empty name")
	}
	
	
	// if name is received, return a value that embeds the name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}