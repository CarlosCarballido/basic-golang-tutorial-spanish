# 01. Introduccion a Go

Go, tambien llamado Golang, es un lenguaje compilado disenado por Google para escribir software simple, rapido y facil de mantener. Su filosofia es opuesta a los lenguajes llenos de abstracciones complejas: Go prefiere pocas reglas, una sintaxis pequena y un ecosistema de herramientas muy consistente.

## Que aporta Go

- Compilacion rapida.
- Codigo legible y predecible.
- Concurrencia nativa con `goroutines` y `channels`.
- Herramientas integradas: `go run`, `go test`, `go fmt`, `go build`.

## Analogia intuitiva

Piensa en Go como un taller con pocas herramientas, pero muy bien elegidas. En Java muchas tareas pueden resolverse de varias maneras. En Go normalmente existe una manera clara y directa de hacerlo, lo que reduce discusiones y errores.

## Primer programa

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola, Go")
}
```

### Explicacion linea por linea

- `package main`: indica que este archivo pertenece al paquete ejecutable principal.
- `import "fmt"`: importa la libreria para salida por consola.
- `func main()`: punto de entrada del programa.
- `fmt.Println(...)`: imprime texto y agrega salto de linea.

## Flujo de trabajo tipico

```text
editar -> go fmt -> go test -> go run o go build
```

## Comparacion con Java

- En Java necesitas una clase con `public static void main(String[] args)`.
- En Go basta con una funcion `main` en el paquete `main`.
- En Java suele haber mas estructura ceremonial; en Go el arranque es mas directo.

## Ejemplo realista

```go
package main

import "fmt"

func main() {
    nombre := "Ana"
    edad := 21
    fmt.Printf("%s tiene %d anos\n", nombre, edad)
}
```

## Errores frecuentes

- Olvidar el `package main` cuando quieres ejecutar el archivo.
- Confundir `fmt.Println` con `fmt.Printf`.
- Intentar usar Go como si fuera Java y buscar demasiada verbosidad.

## Buenas practicas

- Ejecuta `go fmt` antes de entregar codigo.
- Usa nombres cortos y descriptivos.
- Mantiene funciones pequenas y con una responsabilidad clara.

## Preguntas de autoevaluacion

1. Que diferencia hay entre `go run` y `go build`?
2. Por que Go pone tanta importancia en la legibilidad?
3. Que papel cumple el paquete `main`?

## Ejercicio guiado

Escribe un programa que imprima tu nombre, tu carrera y tu ciudad usando `fmt.Printf`.

## Ejercicio sin resolver

Crea un programa que imprima tres lineas: una presentacion, una meta personal y una razon por la que quieres aprender Go.
