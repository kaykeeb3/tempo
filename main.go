package main

import (
	"fmt"
	"math"
)

func tempoSobrevivencia(temperaturaAgua float64) float64 {
	temperaturaCorpo := 37.0 // temperatura normal do corpo em graus Celsius
	const taxaResfriamento = 0.25

	tempo := (temperaturaCorpo - temperaturaAgua) / taxaResfriamento
	return math.Round(tempo)
}

func main() {
	var temperaturaAgua float64

	fmt.Println("Bem-vindo(a) à estimativa de tempo de sobrevivência no mar!")
	fmt.Println("Por favor, informe a temperatura da água em graus Celsius:")

	fmt.Scanln(&temperaturaAgua)

	tempo := tempoSobrevivencia(temperaturaAgua)

	if tempo > 0 {
		fmt.Printf("A estimativa de tempo de sobrevivência é de aproximadamente %.0f minutos.\n", tempo)
	} else {
		fmt.Println("A temperatura da água é maior ou igual à temperatura normal do corpo.")
		fmt.Println("Nesse caso, a estimativa não se aplica, pois o corpo não está em risco de resfriamento excessivo.")
	}
}
