package main

import "fmt"

func IsUpper(s string) bool {
	runes := []rune(s)
	for index := range runes {
		if (runes[index] >= 0 && runes[index] <= 64) || (runes[index] >= 91 && runes[index] <= 127) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}
