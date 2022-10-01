package main

import "fmt"

func ConcatParams(args []string) string {
	n := ""
	for _, letter := range args {
		n = n + letter + "\n"
	}
	return n

}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
