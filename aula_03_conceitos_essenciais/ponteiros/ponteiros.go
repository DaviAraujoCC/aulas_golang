package main

import (
	"fmt"
)

var (
	teste string
)

func main() {

	/**

	Ponteiro é uma variável que armazena um endereço de memória de outra variável.
	Ponteiros são usados para passar referências de valores de um tipo primitivo para funções.

	Para acessar o endereço de memória de uma variável específica, utilizamos o símbolo &.

	Use o * para realizar a desreferenciação no ponteiro (acessar o conteúdo da variável apontada).

	Nota: é muito importante saber aonde vai ser utilizado o ponteiro, visto que dependendo da sua aplicação, o acesso simultaneo ao mesmo endereço de memória pode gerar deadlock e outros erros, o que pode ser resolvido com os pacotes de concorrência.
	**/

	var conta int = 100
	Sacar(&conta, 50)
	fmt.Println(conta)

}

func Sacar(conta *int, valor int) {
	*conta = *conta - valor
}
