# 03. Tipos, variables y constantes

Go tiene un sistema de tipos estatica y fuertemente tipado. Eso significa que el compilador verifica tipos antes de ejecutar el programa, y que no convierte valores de forma magica como otros lenguajes dinamicos.

## Variables

```go
var nombre string = "Lucia"
edad := 20
```

### Explicacion

- `var nombre string = "Lucia"`: declaracion explicita con tipo.
- `edad := 20`: inferencia de tipo en una declaracion corta.

## Constantes

```go
const pi = 3.1416
const mensaje = "Bienvenido"
```

Las constantes no cambian durante la ejecucion. Son utiles para valores fijos, configuraciones y limites.

## Tipos basicos

- Numericos: `int`, `int64`, `float64`, `uint`.
- Booleanos: `bool`.
- Texto: `string`.
- Caracteres Unicode: `rune`.

## Analogia intuitiva

Una variable es una caja etiquetada. El tipo dice que tipo de objeto puede guardar esa caja. Go no deja meter cualquier cosa si la etiqueta dice otra.

## Ejemplo realista

```go
package main

import "fmt"

func main() {
    const tasaIVA = 0.21
    producto := "Teclado"
    precio := 50.0
    total := precio + precio*tasaIVA

    fmt.Printf("Producto: %s\n", producto)
    fmt.Printf("Total con IVA: %.2f\n", total)
}
```

## Conversion explicita

```go
var a int = 10
var b float64 = float64(a)
```

En Go la conversion se escribe de forma explicita. No se asume automaticamente.

## Comparacion con Java

- Go permite `:=` para inferir tipos en variables locales.
- Java usa tipos declarados de forma mas repetitiva.
- Go obliga a pensar antes en el tipo correcto, lo que reduce sorpresas.

## Errores frecuentes

- Confundir asignacion con declaracion.
- Intentar sumar `int` y `float64` sin conversion.
- Pensar que una constante puede modificarse despues.

## Buenas practicas

- Usa `const` cuando el valor no cambie.
- Elige el tipo mas simple que resuelva el problema.
- Evita tipos demasiado grandes si no los necesitas.

## Preguntas de autoevaluacion

1. Que diferencia hay entre `var` y `:=`?
2. Por que Go exige conversion explicita entre tipos numericos?
3. Cuando usarias `const` en un proyecto real?

## Ejercicio guiado

Declara una variable para el nombre de un curso, otra para el numero de creditos y una constante para la nota maxima.

## Ejercicio sin resolver

Escribe un programa que calcule el area de un circulo usando una constante `pi` y una variable de radio.
