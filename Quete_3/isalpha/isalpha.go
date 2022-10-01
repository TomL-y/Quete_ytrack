package main

import "fmt"

func IsAlpha(s string) bool {
	runes := []rune(s)
	for index := range runes {
		if (runes[index] >= 0 && runes[index] <= 47) || (runes[index] >= 58 && runes[index] <= 63) || (runes[index] >= 91 && runes[index] <= 96) || (runes[index] >= 123 && runes[index] <= 127) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you ?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4 ?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
