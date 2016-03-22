// Echo index and arg 1 per line
package main

import (
	"os"
	"fmt"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}