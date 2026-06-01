# 06. Punteros

Los punteros permiten referenciar una direccion de memoria en lugar de copiar un valor. Esto es importante para modificar datos desde funciones y para evitar copias innecesarias en estructuras grandes.

## Sintaxis esencial

```go
edad := 20
p := &edad
fmt.Println(*p)
```

## Analogia intuitiva

Un puntero es como la direccion de una casa. En vez de llevar siempre la casa entera, llevas la direccion y vas alli cuando necesitas acceder a ella.

## Ejemplo completo

```go
package main

import "fmt"

func aumentar(x *int) {
    *x = *x + 1
}

func main() {
    valor := 10
    aumentar(&valor)
    fmt.Println(valor)
}
```

### Linea por linea

- `func aumentar(x *int)`: recibe un puntero a entero.
- `*x = *x + 1`: desreferencia y modifica el entero original.
- `aumentar(&valor)`: pasa la direccion de `valor`.

## Cuidados con nil

```go
var p *int
if p != nil {
    fmt.Println(*p)
}
```

## Comparacion con Java

- Java no expone punteros de forma explicita al programador.
- Go si los expone, pero de forma simple y controlada.
- En Go el objetivo no es manipular memoria manualmente, sino expresar intenciones claras.

## Errores frecuentes

- Desreferenciar un puntero `nil`.
- Confundir copia de valor con referencia.
- Usar punteros sin necesidad y complicar la lectura.

## Buenas practicas

- Usa punteros cuando necesitas mutacion o evitar copias costosas.
- No uses punteros solo por costumbre.
- Comprueba `nil` cuando el puntero pueda no estar inicializado.

## Preguntas de autoevaluacion

1. Que diferencia hay entre `valor` y `&valor`?
2. Cuando conviene pasar un puntero a una funcion?
3. Por que un puntero `nil` puede causar un panic?

## Ejercicio guiado

Escribe una funcion que reciba un puntero a `int` y lo duplique.

## Ejercicio sin resolver

Crea una funcion que reciba un `struct` por puntero y modifique uno de sus campos.
