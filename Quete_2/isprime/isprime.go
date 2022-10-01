package main

import "fmt"

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	var compteur int = 0
	for i := 2; i < nb-1; i++ {
		if nb%i == 0 {
			compteur = compteur + 1
		}
	}
	if compteur != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
}
