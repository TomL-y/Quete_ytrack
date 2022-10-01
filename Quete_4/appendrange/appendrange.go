package main

import "fmt"

func AppendRange(min, max int) []int {
	var reponse []int
	if max >= min {
		for i := min; i < max; i++ {
			reponse = append(reponse, i)
		}
	}
	return reponse
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
