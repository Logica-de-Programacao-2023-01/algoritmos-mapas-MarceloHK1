package main

import "fmt"

func cinco(a string) map[string]int {
	mapa := make(map[string]int)
	for _, letras := range a {
		mapa[string(letras)]++
	}
	return mapa
}
func main() {
	var x string
	fmt.Print("digite uma palavra: ")
	fmt.Scan(&x)

	final := cinco(x)
	fmt.Print(final)
}
