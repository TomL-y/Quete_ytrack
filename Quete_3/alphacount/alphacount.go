package main

import "fmt"

func AlphaCount(s string) int {
	compteur := 0
	for _, letter := range s {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			compteur++
		}
	}
	return compteur
}

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}
