package main

import "fmt"

type Usuario struct {
	Nombre string
	Edad   int
}

func main() {
	usuario := Usuario{Nombre: "Ana", Edad: 21}
	fmt.Printf("%s tiene %d anos\n", usuario.Nombre, usuario.Edad)
}
