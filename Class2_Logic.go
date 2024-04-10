package main

import "fmt"

func main() {
	var nomeusuario = ""
	fmt.Println("Qual é o seu nome")
	fmt.Scanln(&nomeusuario)
	fmt.Println("Welcome " + nomeusuario)

}
