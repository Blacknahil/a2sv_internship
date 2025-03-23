package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered channels"
	messages <- "practice by nahom"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// fmt.Println(<-messages)
	// doing another reciving causes
	// fatal error: all goroutines are asleep - deadlock!
}
