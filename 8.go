package main

import "fmt"

func oito(a map[string]float64) map[string]float64 {
	despesas := make(map[string]float64)

	soma := 0.0
	for _, quantia := range a {
		soma += quantia
	}
	for nomes, quantia := range a {
		media := soma / float64(len(a))
		despesas[nomes] = quantia - media
	}
	return despesas
}
func main() {
	x := map[string]float64{
		"a": 10.0,
		"b": 5.0,
		"c": 15.0,
	}
	final := oito(x)
	fmt.Print(final)
}
