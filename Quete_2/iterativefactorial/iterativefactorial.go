package main

import "fmt"

func IterativeFactorial(nb int) int {
	x := 1
	for i := 1; i <= nb; i++ {
		x = x * i
	}
	return x
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
