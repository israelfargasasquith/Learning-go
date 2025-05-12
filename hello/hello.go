package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"

	"log"
)

func main() {
	fmt.Println(quote.Go())

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{"Raeliss", "Dante", "Vergil"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
