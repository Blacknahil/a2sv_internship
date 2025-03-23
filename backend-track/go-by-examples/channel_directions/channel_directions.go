package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

// it recives a message through a channel(pings) and sends that message though a channel(pongs)
func pongs(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	pings_channel := make(chan string, 1)
	pongs_channel := make(chan string, 1)

	ping(pings_channel, "passed message")
	pongs(pings_channel, pongs_channel)

	fmt.Println(<-pongs_channel)
}
