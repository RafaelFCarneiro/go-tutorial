package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func Hello(names []string) map[string]string {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	msg, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	return msg
}

func main() {
	fmt.Println(Hello([]string{"Kent", "Eric", "Vernon"}))
}
