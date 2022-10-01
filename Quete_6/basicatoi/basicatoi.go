package main

import "fmt"

func BasicAtoi(s string) int {
	var tab int
	runes := []rune(s)
	for n := range runes {
		if runes[n] >= 49 && runes[n] <= 57 {
			tab = tab*10 + int(runes[n]-48)
		}
	}
	return tab
}
func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
