package main

import (
	"fmt"
	"strconv"
)

func main() {

	mensagem := "Olá, mundo!"

	resultado, n := mostrarMensagem(mensagem, 10)

	fmt.Println(resultado + ",repetido " + strconv.Itoa(n))

}

func mostrarMensagem(mensagem string, n int) (string, int) {
	for i := 0; i < n; i++ {
		fmt.Println(mensagem)
	}

	return "Finalizou o programa", n
}
