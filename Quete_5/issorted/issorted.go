package main

import "fmt"

func IsSorted(f func(a, b int) int, tab []int) bool {
	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			return false
		}
	}
	return true
}

func Sorted(a int, b int) int {
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(Sorted, a1)
	result2 := IsSorted(Sorted, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
