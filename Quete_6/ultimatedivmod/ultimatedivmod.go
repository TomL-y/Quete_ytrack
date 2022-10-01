package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	var div int
	var mod int
	div = *a / *b
	mod = *a % *b
	*a = div
	*b = mod
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
