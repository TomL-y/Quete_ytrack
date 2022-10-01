package main

import "fmt"

func Join(strs []string, sep string) string {
	n := ""
	for i, letter := range strs {
		if i == len(strs)-1 {
			n = n + letter
		} else {
			n = n + letter + sep
		}
	}
	return n
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you ?"}
	fmt.Println(Join(toConcat, ":"))
}
