package main

import (
	"fmt"
)

func phi(n int) int {
	res := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			res -= res / i
		}
	}
	if n > 1 {
		res -= res / n
	}
	return res
}

func main() {
	fmt.Println(phi(21))
	fmt.Println(phi(7))
}
