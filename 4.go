package main

import (
	"sort"
	"strings"
)

func quatro(a []string) map[string][]string {
	mapa := make(map[string][]string)

	for _, palavra := range a {
		anagrama := anagramas(strings.ToLower(palavra))
		if _, ok := mapa[anagrama]; !ok {
			mapa[anagrama] = []string{palavra}
		} else {
			mapa[anagrama] = append(mapa[anagrama], palavra)
		}
	}
	return mapa
}
func anagramas(a string) string {
	letras := strings.Split(a, "")
	sort.Strings(letras)
	return strings.Join(letras, "")
}
