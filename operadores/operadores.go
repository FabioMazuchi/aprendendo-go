package main

import "fmt"

func main() {
	x, y := 10, 5
	soma := x + y
	subtracao := x - y
	divisao := x / y
	modulo := x % y
	multiplicacao := x * y

	fmt.Printf("%v + %v = %v\n", x, y, soma)
	fmt.Printf("%v - %v = %v\n", x, y, subtracao)
	fmt.Printf("%v / %v = %v\n", x, y, divisao)
	fmt.Println(x, "%", y, "=", modulo)
	fmt.Printf("%v x %v = %v\n", x, y, multiplicacao)
}
