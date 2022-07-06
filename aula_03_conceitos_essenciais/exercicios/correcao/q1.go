package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	pessoa := map[string]string{}

	pessoa["nome"] = "davi"
	pessoa["idade"] = "23"
	pessoa["estado"] = "ceara"

	fmt.Println(pessoa)

	bytesJson, err := json.Marshal(pessoa)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytesJson))

}
