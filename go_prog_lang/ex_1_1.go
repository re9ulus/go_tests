// ex1.1, modify echo program to also print args[0]
package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}