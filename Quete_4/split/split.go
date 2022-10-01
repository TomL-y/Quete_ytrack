package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	var tab []string
	index := -1
	for i := 0; len(s) >= 1; i++ {
		index++
		if index == len(s) {
			tab = append(tab, s)
			break
		}
		if rune(s[index]) == rune(sep[0]) {
			if s[index:index+len(sep)] == sep {
				tab = append(tab, s[:index])
				s = s[index+len(sep):]
				index = -1 //remettre l'index à 0 pour la prochaine itération
			}
		}
	}
	return tab
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
