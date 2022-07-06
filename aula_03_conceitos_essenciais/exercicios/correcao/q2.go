package main

import "fmt"

func main() {

	var v1 = 10
	var v2 = 35

	if ParOuImpar(v1, v2) {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é impar")
	}

}

func ParOuImpar(v1, v2 int) bool {

	total := v1 + v2

	return total%2 == 0
}
