package main

import "fmt"

func SortWordArr(tab []string) {
	for i := 0; i < len(tab)-1; i++ {
		if tab[i] > tab[i+1] {
			tab[i], tab[i+1] = tab[i+1], tab[i]
			SortWordArr(tab)
		}
	}
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println(result)
}
