package main

import "fmt"

func Atoi(s string) int {
	var tab int
	count := 0
	for _, n := range s {
		if n >= 49 && n <= 57 || n == 45 || n == 43 {
			if n == 45 || n == 43 {
				count += 1
			} else {
				tab = tab*10 + int(n-48)
			}
		}
		if n == 32 {
			return 0
		}
	}
	if count == 1 && string(s[0]) == "-" {
		tab = tab - tab*2
	} else if count == 2 {
		return 0
	}
	count = 0
	return tab
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
