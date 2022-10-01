package main

import "fmt"

func Any(f func(string) bool, a []string) bool {
	for _, n := range a {
		if f(n) == true {
			return true
		}
	}
	return false
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
	a1 := []string{"This", "is", "4", "you"}
	a2 := []string{"Hello", "how", "are", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
