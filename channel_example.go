package main

import (
	"fmt"
	"log"
)

//START OMIT
func Greet(name string, response_chan chan string) {
	greeting := fmt.Sprintf("Greetings, %s!", name)
	response_chan <- greeting
}

func main() {

	cs := make(chan string)
	go Greet("Alice", cs)
	greeting := <-cs
	log.Print(greeting)
}

//END OMIT
