package main

import "fmt"

func dois(a map[string]int, b map[string]int) map[string]int {
	mapa := make(map[string]int)
	for nome, posiçao := range a {
		mapa[nome] = posiçao
	}
	for mes, dia := range b {
		mapa[mes] = dia
	}
	return mapa
}
func main() {
	id := map[string]int{
		"marcelo": 1,
		"caio":    2,
	}
	data := map[string]int{
		"novembro": 29,
		"julho":    27,
	}
	final := dois(id, data)
	fmt.Print(final)
}
