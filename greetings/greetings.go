package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a greeting for every named person
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// go executes init functions automaticaly after global vars have been initialized
func init() {
	rand.Seed(time.Now().UnixNano())
}

// lowercase letter for function == private function reachable only from its own package
func randomFormat() string {
	formats := []string{ // slice
		"Hi, %v. Welcome!",
		"Hello, %v",
		"FYI, %v",
	}

	return formats[rand.Intn(len(formats))]
}
