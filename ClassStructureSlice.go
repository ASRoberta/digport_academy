package main

import "fmt"

func main() {

	const oi = "Hello"

	s := make([]string, 0, 4)
	s = append(s, oi, "Maria", "Ana", "Julia", "Rafaela")
	v := s[2]

	fmt.Println("esste é o meu slice:", s)
	fmt.Println(s[0], v)

}
