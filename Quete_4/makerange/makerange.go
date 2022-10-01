package main

import "fmt"

func MakeRange(min, max int) []int {
	size := max - min
	if size >= 0 {
		reponse := make([]int, size)
		for i := min; i < max; i++ {
			reponse[i-size] = i
		}
		return reponse
	}
	var vide []int
	return vide
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
