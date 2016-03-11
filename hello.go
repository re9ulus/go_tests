package main

import (
	"fmt"
)

func SayHello(name string) {
	fmt.Println("Hello, " + name + "!")
}

func main() {
	SayHello("Go")
}
