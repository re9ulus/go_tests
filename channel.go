package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "Hello"
	messages <- "this is"
	go func() {messages <- "channel test!"}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}