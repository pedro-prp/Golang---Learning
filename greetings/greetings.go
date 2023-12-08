package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message

	// if name is empty
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name is received, return a value that embeds the name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

// Hellos returns a map that associates each of the named people
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)

	// For loop to receive the slice of names
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// Associate in the map message with name
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by specifying
	// a random index for the slice of formats

	return formats[rand.Intn(len(formats))]
}
