package main

import "fmt"

func main() {

	slice := []int{}

	slice = append(slice, 0)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)

	fmt.Println(slice)

}
