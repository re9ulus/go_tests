package main

import (
	"fmt"
	"os/exec"
)

func cmdCall(command string, params ...string) (string, error) {
	fmt.Println(command, params)
	c := exec.Command(command, params...)
	res, err := c.Output()
	fmt.Println(string(res))
	return string(res), err
}

func main() {
	if response, err := cmdCall("git", "--help"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(">> Start output: ")
		fmt.Println(response)
		fmt.Println(">> End output: ")
	}
}
