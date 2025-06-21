package main

import (
	"fmt"
	"time"
)

func main() {
	// if use go it mean we want to work parallel

	//go greet("Nice to meet you!")
	//go greet("How are you?")

	done := make(chan bool)
	go slowGreet("How...are...you...?", done)

	//go greet("T hope you're liking the course!")
	<-done
}

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true
}
