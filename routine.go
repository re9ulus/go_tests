package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello, " + name)
}

func main() {
	names := []string{"World", "Go", "All", "Someone"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
		go sayHello(names[i])
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}