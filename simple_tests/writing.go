package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	text := []byte("hello\nworld\n")
	err := ioutil.WriteFile("./test.txt", text)
	check(err)
}