package main

func seis(a []map[string]int) map[string]int {
	mapa := make(map[string]int)

	for _, contagem := range a {
		for palavra, quantidade := range contagem {
			mapa[palavra] += quantidade
		}
	}
	return mapa
}
