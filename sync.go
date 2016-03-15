package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("Worker done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<- done // Blocking operation

	fmt.Println("Main done")
}