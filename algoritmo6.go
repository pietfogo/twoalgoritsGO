package main

import "fmt"

func main() {
	var recebeVal int
	fatorial := 1

	fmt.Print("Digite um número e eu calcularei seu fatorial: ")
	fmt.Scan(&recebeVal)

	for i := 1; i <= recebeVal; i++ {
		fatorial *= i
	}

	fmt.Printf("O valor do fatorial é: %v", fatorial)
}
