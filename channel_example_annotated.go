package main

import (
	"fmt"
	"log"
	"time"
)

//START OMIT
func Greet(name string, response_chan chan string) {
	time.Sleep(3 * time.Second)
	greeting := fmt.Sprintf("Greetings, %s!", name)
	response_chan <- greeting
	log.Print("Greeting function is tired - sleeping for a bit")
	time.Sleep(3 * time.Second)
	log.Print("now the greeting function is done sleeping - terminating")
}

func main() {

	cs := make(chan string)
	go Greet("Alice", cs)
	log.Print("Continuing execution while we wait for greeter to respond")
	greeting := <-cs
	log.Print(greeting)
	time.Sleep(10 * time.Second)
}

//END OMIT
