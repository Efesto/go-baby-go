package main

import (
	"fmt"
	"log"

	"whatever.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("peterson")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	messages, err := greetings.Hellos([]string{"peterson", "anderson", "Manoj"})

	if err != nil {
		log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
	}

}
