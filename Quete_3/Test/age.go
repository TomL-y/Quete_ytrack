package main

import "github.com/01-edu/z01"

func CheckAge(n) {
	if n >= 18 {
		return z01.PrintRune("Tu peux rentrer")
	}
	return z01.PrintRune("Sortez !")
}

func maain() {
	z01.PrintRune(CheckAge(17))
}
