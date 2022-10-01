package main

func PrintWordsTables(a []string) {
	for _, letter := range a {
		print(letter, "\n")
	}
}

func main() {
	a := []string{"Hello", "how", "are", "you?"}
	PrintWordsTables(a)
}
