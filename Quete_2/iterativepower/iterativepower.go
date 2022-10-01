package main

import (
	"fmt"
)

func IterativePower(nb int, power int) int {
	result := 1
	if power <= 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	for i := 1; i <= power; i++ {
		result = result * nb
	}
	return result
}

func main() {
	fmt.Println(IterativePower(10, 5))
}
