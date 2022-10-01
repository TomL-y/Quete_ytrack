package main

import "fmt"

func ToLower(s string) string {
	var runes string
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			runes += string(letter + 32)
		} else {
			runes += string(letter)
		}
	}
	return runes
}

func main() {
	fmt.Println(ToLower("HELLO ! HOW ARE YOU ?"))
}
