package main

import (
	"fmt"
	"strings"
)

func sete(a string) map[string]int {
	mapa := make(map[string]int)
	palavras := strings.Fields(a)

	for _, palavra := range palavras {
		mapa[palavra] = len(palavra)
	}
	return mapa
}
func main() {
	x := "marcelo honda"
	final := sete(x)
	fmt.Print(final)
}
