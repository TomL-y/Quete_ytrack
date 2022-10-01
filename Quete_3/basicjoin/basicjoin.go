package main

import "fmt"

func BasicJoin(elems []string) string {
	n := ""
	for _, letter := range elems {
		n = n + letter
	}
	return n
}

func main() {
	elems := []string{"Hello!", " How", " are", " you ?"}
	fmt.Println(BasicJoin(elems))
}
