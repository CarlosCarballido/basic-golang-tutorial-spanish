# 02. Sintaxis basica

La sintaxis de Go es pequena, pero tiene decisiones importantes: los bloques usan llaves, no hay punto y coma en el codigo normal, y la estructura general del archivo es predecible.

## Elementos esenciales

- `package`
- `import`
- `func`
- `if`, `for`, `switch`
- comentarios `//` y `/* */`

## Analogia intuitiva

Go funciona como una receta breve: primero defines el contexto, luego importas lo que necesitas y despues escribes acciones claras. Si una receta tiene demasiados pasos innecesarios, Go la considera ruido.

## Ejemplo minimo

```go
package main

import "fmt"

func main() {
    if true {
        fmt.Println("Siempre se ejecuta")
    }
}
```

### Linea por linea

- `if true`: condicion booleana.
- Las llaves delimitan el bloque que se ejecuta.
- No se necesita `else` si no hay rama alternativa.

## Bucles en Go

Go no tiene `while`. Usa `for` para todo.

```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}

for condicion {
    // equivalente a while
}

for {
    // bucle infinito
}
```

## Switch

```go
dia := 3
switch dia {
case 1:
    fmt.Println("Lunes")
case 2:
    fmt.Println("Martes")
default:
    fmt.Println("Otro dia")
}
```

## Comparacion con Java

- `for` de Go puede reemplazar `while` y `do-while`.
- `switch` es mas flexible y no exige `break` al final de cada caso.
- Los bloques son simples y la sintaxis evita redundancia.

## Errores frecuentes

- Poner `break` innecesarios al estilo Java.
- Separar mal las llaves y perder legibilidad.
- Olvidar que `for` es la unica estructura de iteracion general.

## Buenas practicas

- Prefiere condiciones simples.
- Usa `switch` cuando hay varios casos claros.
- No abuses de `if` anidados cuando un `switch` simplifica la lectura.

## Preguntas de autoevaluacion

1. Por que Go usa `for` para casi todos los bucles?
2. Que diferencia hay entre `switch` en Go y en Java?
3. Cuando es mas legible un `switch` que una cadena de `if`?

## Ejercicio guiado

Escribe un programa que recorra del 1 al 5 e imprima si cada numero es par o impar.

## Ejercicio sin resolver

Crea un programa que lea una variable `nota` y muestre una calificacion textual usando `switch`.
