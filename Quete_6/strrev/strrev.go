package main

import "fmt"

func StrRev(s string) string {
	var tab string
	for _, n := range s {
		tab = string(n) + tab
	}
	return tab

}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
