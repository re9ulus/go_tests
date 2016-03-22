package main

import "fmt"

// chan<- only for writing
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// <-chan only for reading
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg + "?"
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "This is a question")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
