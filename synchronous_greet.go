package main

import (
	"log"
	"time"
)

//START OMIT
func Greet(name string) {
	log.Printf("Greetings, %s!", name)
	time.Sleep(3 * time.Second)
}

func main() {
	Greet("Alice")
	Greet("Bob")

	time.Sleep(5 * time.Second)
}

//END OMIT
