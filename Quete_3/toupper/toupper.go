package main

import "fmt"

func ToUpper(s string) string {
	var runes string
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			runes += string(letter - 32)
		} else {
			runes += string(letter)
		}
	}
	return runes
}

func main() {
	fmt.Println(ToUpper("Hello! How are you ?"))
}
