package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	EmptyName = "empty name"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New(EmptyName)
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Returns named people with greeting message
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	// standard go for loop
	// range returns element and its index
	// index is ignored using blank identifier _
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// add name as a key and the message as a value to messages map
		messages[name] = message
	}
	return messages, nil
}

// Go executes init functions automatically at program startup,
// after global variables have been initialized.
// sets initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns selected random greeting message
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v Well met!",
	}

	// Return randomly selected format by specifying random index
	return formats[rand.Intn(len(formats))]
}
