# 05. Structs y composicion

Los `structs` agrupan datos relacionados. En Go no se piensa tanto en herencia como en composicion: en lugar de decir "es un tipo de", Go prefiere "tiene un tipo de" o "incluye un tipo".

## Struct basico

```go
type Usuario struct {
    Nombre string
    Edad   int
}
```

## Instanciacion

```go
u := Usuario{Nombre: "Marta", Edad: 23}
```

## Composicion por embebido

```go
type Persona struct {
    Nombre string
}

type Alumno struct {
    Persona
    Curso string
}
```

## Analogia intuitiva

Un struct es como una ficha de datos. La composicion es como meter una ficha dentro de otra para reutilizar atributos sin construir jerarquias pesadas.

## Ejemplo realista

```go
package main

import "fmt"

type Producto struct {
    Nombre string
    Precio float64
}

func (p Producto) ConIVA() float64 {
    return p.Precio * 1.21
}

func main() {
    p := Producto{Nombre: "Monitor", Precio: 200}
    fmt.Printf("%s cuesta %.2f con IVA\n", p.Nombre, p.ConIVA())
}
```

### Linea por linea

- `type Producto struct`: define un nuevo tipo compuesto.
- `func (p Producto) ConIVA()`: metodo asociado al struct.
- `p.Precio * 1.21`: calcula el precio final.

## Comparacion con Java

- En Java solemos crear clases con campos privados y getters/setters.
- En Go se usan structs simples y se decide con cuidado que hace publico cada campo.
- La composicion en Go suele ser mas natural que una jerarquia de clases profunda.

## Errores frecuentes

- Escribir nombres de campos sin pensar en si deben exportarse o no.
- Mezclar demasiados conceptos en un solo struct.
- Abusar de la composicion sin claridad semantica.

## Buenas practicas

- Mantiene structs pequenos y coherentes.
- Usa composicion cuando aporta reutilizacion real.
- Prefiere nombres claros antes que abreviaturas confusas.

## Preguntas de autoevaluacion

1. Que ventaja tiene la composicion frente a la herencia?
2. Cuando un struct esta demasiado cargado?
3. Que diferencia hay entre campos exportados y no exportados?

## Ejercicio guiado

Define un struct `Libro` con titulo, autor y paginas. Crea una funcion que devuelva el resumen del libro.

## Ejercicio sin resolver

Modela un `Pedido` con un cliente, una lista de productos y un total calculado.
