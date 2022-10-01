package main

import "fmt"

func IsLower(s string) bool {
	runes := []rune(s)
	for index := range runes {
		if (runes[index] >= 0 && runes[index] <= 96) || (runes[index] >= 123 && runes[index] <= 127) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}
