package main

import (
	"fmt"
	"strings"
)

func um(texto string) map[string]int {
	palavras := strings.Fields(texto)
	mapa := make(map[string]int)

	for _, palavra := range palavras {
		mapa[palavra]++
	}
	return mapa
}
func main() {
	x := "a b a b a c c"
	final := um(x)
	fmt.Print(final)
}
