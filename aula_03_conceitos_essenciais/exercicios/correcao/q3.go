package main

import "fmt"

func main() {

	var conta = 100
	fmt.Println("Antes do saque:")
	fmt.Println(conta)

	Sacar(&conta, 50)
	fmt.Println("Depois do saque:")
	fmt.Println(conta)

	Depositar(&conta, 60)
	fmt.Println("Depois do deposito:")
	fmt.Println(conta)

}

func Sacar(conta *int, valor int) {
	*conta = *conta - valor
}

func Depositar(conta *int, valor int) {
	*conta = *conta + valor
}
