package main

import "fmt"

func main() {
	var numeros = [4]int{1, 2, 3, 4}
	produto := 1
	for i := 0; i < len(numeros); i++ {
		produto *= numeros[i]
	}
	fmt.Print(produto)
}
