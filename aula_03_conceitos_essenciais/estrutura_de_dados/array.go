package main

import (
	"fmt"
	"strconv"
)

func main() {

	var array [5]int

	for i := 0; i < len(array); i++ {
		array[i] = i * 10
	}

	for posicao, value := range array {
		fmt.Println("posicao " + strconv.Itoa(posicao))
		fmt.Println("valor " + strconv.Itoa(value))
		fmt.Println("")
	}

}
