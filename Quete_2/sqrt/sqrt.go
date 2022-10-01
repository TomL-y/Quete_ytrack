package main

import "fmt"

func main() {
	fmt.Println(Sqrt(225))
}

func Sqrt(nb int) int {
	for i := 1; nb/i >= i; i++ {
		if nb/i == i {
			return i
		}
	}
	return 0
}
