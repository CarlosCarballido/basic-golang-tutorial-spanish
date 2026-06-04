package main

import "fmt"

func dividir(a, b int) int {
	if b == 0 {
		panic("division por cero")
	}
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperado de:", r)
		}
	}()

	fmt.Println(dividir(10, 0))
	fmt.Println("fin normal")
}
