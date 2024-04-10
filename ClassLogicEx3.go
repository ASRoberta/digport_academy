package main

import "fmt"

func main() {

	var numero int
	fmt.Println("Digite um numero inteiro: ")
	fmt.Scanln(&numero)

	if numero > 0 {
		fmt.Println("Você digitou um número positivo")
	} else if numero < 0 {
		fmt.Println("Você digitou um número negativo")
	} else {
		fmt.Println("Você digitou Zero")
	}

}
