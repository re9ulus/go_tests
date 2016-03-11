package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if (e != nil) {
		panic(e)
	}
}

func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return string(data)
}

func main() {
	var fileContent string = ReadFile("./hello.go")
	fmt.Println(fileContent)
}
