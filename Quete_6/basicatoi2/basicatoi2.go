package main

import "fmt"

func BasicAtoi2(s string) int {
	var tab int
	runes := []rune(s)
	for n := range runes {
		if runes[n] >= 49 && runes[n] <= 57 {
			tab = tab*10 + int(runes[n]-48)
		}
		if runes[n] == 32 {
			return 0
		}
	}
	return tab
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
