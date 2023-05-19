package main

import "fmt"

func tres(a map[string]int) int {
	soma := 0
	for _, numeros := range a {
		soma += numeros
	}
	return soma
}
func main() {
	numeros := map[string]int{
		"primeiro": 1,
		"segundo":  2,
		"terceiro": 3,
	}
	final := tres(numeros)
	fmt.Print(final)
}
