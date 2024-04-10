package main

import "fmt"

func main() {
	var numusuario int = 1
	fmt.Println(numusuario)
	var numerousuario int = 2
	fmt.Println("Qual é o seu numero")
	fmt.Scanln(&numerousuario)
	fmt.Println("O resultado é:", numusuario+numerousuario)

}
