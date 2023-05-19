package main

import "fmt"

func nove(a int) map[int]int {
	mapa := make(map[int]int)
	var sequencia = fibonacci(a)
	for posicao, numero := range sequencia {
		mapa[posicao] = numero

	}

	return mapa
}
func fibonacci(a int) []int {
	slice := make([]int, a+1, a+2)

	if a < 2 {
		slice = slice[0:2]
	}
	slice[0] = 0
	slice[1] = 1
	for i := 2; i <= a; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
func main() {
	a := 5
	final := nove(a)
	fmt.Print(final)
}
