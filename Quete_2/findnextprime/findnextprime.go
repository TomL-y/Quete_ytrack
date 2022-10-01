package main

import "fmt"

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 0
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			nb = nb + 1
		}
	}
	return nb
}
func main() {
	fmt.Println(FindNextPrime(6))
	fmt.Println(FindNextPrime(4))
}
