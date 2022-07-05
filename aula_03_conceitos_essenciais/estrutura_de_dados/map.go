package main

import "fmt"

func main() {

	maptux := map[string]string{}

	var maptest = make(map[string]string)

	maptest["nome"] = "tux"

	maptux["nome"] = "tux"
	maptux["idade"] = "18"

	fmt.Println(maptux)

	fmt.Println(maptux["nome"])
	fmt.Println(maptest["nome"])

}
