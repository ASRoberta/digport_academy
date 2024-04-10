package main

import "fmt"

type Funcionario struct {
	nome    string
	idade   int
	role    string
	salario float64
}

func main() {

	var arr [3]Funcionario
	Funcionario0 := Funcionario{nome: "Rô", idade: 31, role: "Manager", salario: 2.9}
	Funcionario1 := Funcionario{nome: "Su", idade: 30, role: "Economist", salario: 3.9}
	Funcionario2 := Funcionario{nome: "Vi", idade: 28, role: "Accountant", salario: 4.9}

	arr[0] = Funcionario0
	arr[1] = Funcionario1
	arr[2] = Funcionario2

	fmt.Println(arr[1])

}
