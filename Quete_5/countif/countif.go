package main

import "fmt"

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, n := range tab {
		if f(n) == true {
			count += 1
		}
	}
	return count
}

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
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
