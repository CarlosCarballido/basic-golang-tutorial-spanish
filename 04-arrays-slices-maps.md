# 04. Arrays, slices y maps

Estos tres tipos cubren colecciones, pero resuelven problemas distintos. Un array tiene tamano fijo. Un slice es una vista flexible sobre una secuencia. Un map es una tabla clave-valor.

## Arrays

```go
var notas [3]int
notas[0] = 7
notas[1] = 8
notas[2] = 10
```

El tamano forma parte del tipo: `[3]int` no es lo mismo que `[5]int`.

## Slices

```go
numeros := []int{1, 2, 3}
numeros = append(numeros, 4)
```

Un slice es mas util en la practica porque puede crecer y encadena operaciones de forma sencilla.

## Maps

```go
edades := map[string]int{
    "Ana": 20,
    "Luis": 22,
}
```

## Analogia intuitiva

- Array: estanteria con huecos fijos.
- Slice: lista dinamica que observa o amplía esa estanteria.
- Map: diccionario donde buscas por palabra, no por posicion.

## Diagrama ASCII

```text
array: [10][20][30]
slice: inicio -> [10][20][30]
map:   "Ana" -> 20
       "Luis" -> 22
```

## Ejemplo realista

```go
package main

import "fmt"

func promedio(notas []int) float64 {
    suma := 0
    for _, nota := range notas {
        suma += nota
    }
    return float64(suma) / float64(len(notas))
}

func main() {
    notas := []int{6, 7, 9, 10}
    fmt.Println("Promedio:", promedio(notas))
}
```

### Linea por linea

- `notas []int`: parametro slice de enteros.
- `for _, nota := range notas`: recorre valores ignorando el indice.
- `len(notas)`: longitud del slice.
- `float64(...)`: conversion para evitar division entera.

## Mapas en practica

```go
contador := map[string]int{}
contador["go"]++
contador["go"]++
fmt.Println(contador["go"]) // 2
```

## Comparacion con Java

- Un array de Go es mas estricto que un array Java porque el tamano es parte del tipo.
- Un slice se parece a `ArrayList`, pero con semantica mas ligera.
- Un map se parece a `HashMap`, aunque la API es mas pequena.

## Errores frecuentes

- Confundir array con slice.
- Olvidar inicializar un map antes de escribir en el.
- Usar `append` sin reasignar el resultado.

## Buenas practicas

- Usa slices por defecto y arrays solo cuando el tamano fijo sea importante.
- Comprueba la presencia de una clave en un map con la forma de dos valores.
- Prefiere funciones que reciban slices en vez de arrays si no hay una razon fuerte para fijar el tamano.

## Preguntas de autoevaluacion

1. Que diferencia hay entre `[]int` y `[3]int`?
2. Por que `append` puede devolver un nuevo slice?
3. Cuando conviene usar un `map`?

## Ejercicio guiado

Crea un slice con cinco nombres, agrega uno mas con `append` y recorre la lista imprimiendo cada nombre.

## Ejercicio sin resolver

Escribe un programa que cuente cuantas veces aparece cada letra en una palabra usando un `map[rune]int`.
