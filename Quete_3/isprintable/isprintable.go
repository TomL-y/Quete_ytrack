package main

import "fmt"

func IsPrintable(s string) bool {
	runes := []rune(s)
	for index := range runes {
		if runes[index] >= 0 && runes[index] <= 32 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}
