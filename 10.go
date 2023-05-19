package main

import "fmt"

func dez(a []int) map[[2]int]int {
	mapa := make(map[[2]int]int)

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			pares := [2]int{a[i], a[j]}
			mapa[pares]++
		}
	}
	return mapa
}
func main() {
	slice := []int{1, 2, 3, 2, 3}

	resultado := dez(slice)
	fmt.Println(resultado)
}
