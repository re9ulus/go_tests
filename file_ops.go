package main

import (
	"os"
	"fmt"
	"bufio"
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

func ReadNBytes(filename string, n int) string {
	f, err := os.Open(filename)
	check(err)
	buf := make([]byte, n)
	_, err = f.Read(buf)
	check(err)
	f.Close()
	return string(buf)
}

func ReadFileBuf(filename string, n int) string {
	f, err := os.Open(filename)
	check(err)
	reader := bufio.NewReader(f)
	buf, err := reader.Peek(n)
	check(err)
	f.Close()
	return string(buf)
}


func main() {
	var filename string = "./hello.go"
	fmt.Println(">> Read whole file: ")
	var fileContent string = ReadFile(filename)
	fmt.Println(fileContent)

	var n int = 10

	fmt.Printf(">> Read first %d bytes\n", n)
	fileContent = ReadNBytes(filename, n)
	fmt.Println(fileContent)

	fmt.Printf(">> Read first %d bytes with buffio\n", n)
	fileContent = ReadFileBuf(filename, n)
	fmt.Println(fileContent)
}
