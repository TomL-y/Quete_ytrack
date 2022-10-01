package main

import "fmt"

func IsNumeric(s string) bool {
	runes := []rune(s)
	for index := range runes {
		if (runes[index] >= 0 && runes[index] <= 47) || (runes[index] >= 58 && runes[index] <= 127) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
