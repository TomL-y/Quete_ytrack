package main

import "fmt"

func Map(f func(int) bool, a []int) []bool {
	var tab []bool
	for _, n := range a {
		if f(n) == true {
			tab = append(tab, true)
		} else {
			tab = append(tab, false)
		}
	}
	return tab
}

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
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
