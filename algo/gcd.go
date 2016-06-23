package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	fmt.Println(gcd(2, 7))
	fmt.Println(gcd(6, 27))
	fmt.Println(lcm(2, 6))
}
