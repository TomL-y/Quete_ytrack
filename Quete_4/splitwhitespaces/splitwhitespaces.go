package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	var tab []string
	mot := ""
	for _, letter := range s {
		if string(letter) != " " {
			mot += string(letter)
		} else {
			tab = append(tab, mot)
			mot = ""
		}
	}
	if mot != "" {
		tab = append(tab, mot)
	}
	return tab
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
