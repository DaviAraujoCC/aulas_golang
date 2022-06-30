package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {

	var pessoa Pessoa

	pessoa.Nome = "Davi"
	pessoa.Idade = 22

	bytes, err := json.Marshal(pessoa)
	if err != nil {
		panic(err)
	}

	var pessoaNova Pessoa

	err = json.Unmarshal(bytes, &pessoaNova)
	if err != nil {
		panic(err)
	}

	fmt.Println(pessoaNova.Nome)
	fmt.Println(pessoaNova.Idade)

}
