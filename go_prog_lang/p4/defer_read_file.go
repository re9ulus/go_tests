package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s, err := ReadFile("./append.go")
	if err != nil {
		panic("Error")
	}
	fmt.Println(string(s[:]))
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
