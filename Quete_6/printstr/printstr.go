package main

func PrintStr(s string) {
	for _, c := range s {
		print(string(c))
	}

}

func main() {
	PrintStr("Hello World!")
}
