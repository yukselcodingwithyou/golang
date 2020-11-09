package main

import (
	"fmt"
	"golang/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Error handling:  Return an error as a value so the caller can check for it
	message, err := greetings.Hello("Yuksel")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Ahmet", "Mehmet", "Ali"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
