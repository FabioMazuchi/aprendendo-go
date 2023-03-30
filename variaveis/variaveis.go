package main

import "fmt"

func main() {
	// nome := "Fábio"
	// idade := 39
	// altura := 1.65

	nome, idade, altura := "Fábio", 39, 1.65

	fmt.Println("Meu nome é", nome, ",tenho", idade, "anos e minha altura é", altura)

	fmt.Printf("%T, %T, %T,", nome, idade, altura)
}
