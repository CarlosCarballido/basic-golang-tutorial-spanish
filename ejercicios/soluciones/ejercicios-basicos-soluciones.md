# Ejercicios basicos

## Ejercicio 1
Enunciado: Escribe un programa que muestre tu nombre, tu edad y tu ciudad.
Pistas: usa `fmt.Printf`.
Practica: variables, tipos basicos, salida por consola.


```go

package main

import(
    "fmt"
)

func main(){
    var nombre string = "Carlos"
    var edad  int = 22
    var ciudad string  = "Teo"

    fmt.Printf("%s, %d, %s", nombre, edad, ciudad )
    fmt.Println(nombre, edad, ciudad)
}

```

## Ejercicio 2
Enunciado: Calcula el area de un rectangulo a partir de ancho y alto.
Pistas: define una funcion pequeña.
Practica: variables, operaciones, funciones.


```go

package main

import "fmt"

func calcularArea(x *float64, y *float64) float64{
    return *x * *y
}

func main(){
    x := 2.0
    y := 3.0

    area := calcularArea(&x, &y)

    fmt.Println(area)
}

```

## Ejercicio 3
Enunciado: Imprime los numeros del 1 al 10.
Pistas: usa `for`.
Practica: bucles.


```go

package main

import(
    "fmt"
)

func main(){
    for i:=0 ; i < 10 ; i++{
        fmt.Println(i)
    }
}

```

## Ejercicio 4
Enunciado: Imprime si un numero es par o impar.
Pistas: usa el operador modulo.
Practica: condicionales, bucles.


```go

package main

import(
    "fmt"
)

func esPar(a *int) bool{
    if *a % 2 == 0{
        return true
    } else {
        return false
    }
}

func main(){
    num1 := 1
    num2 := 2

    if esPar(&num1){
        fmt.Println("Par")
    } else{
        fmt.Println("Impar")
    }

    if esPar(&num2){
        fmt.Println("Par")
    } else{
        fmt.Println("Impar")
    }
}

```

## Ejercicio 5
Enunciado: Crea un slice con cinco materias y recorre sus elementos.
Pistas: usa `range`.
Practica: slices, iteracion.


```go

package main

import(
    "fmt"
)

func main(){
    // alternativa: materias := []string{"SIRE", "APAUII", "APAUBIO", "PIC", "EJIA"}
    var materias []string = []string{"SIRE", "APAUII", "APAUBIO", "PIC", "EJIA"}

    for _, materia := range materias{
        fmt.Println(materia)
    }
}

```

## Ejercicio 6
Enunciado: Agrega una materia mas a tu slice anterior.
Pistas: usa `append`.
Practica: slices, crecimiento dinamico.


```go

package main

import(
    "fmt"
)

func main(){
    // alternativa: materias := []string{"SIRE", "APAUII", "APAUBIO", "PIC", "EJIA"}
    var materias []string = []string{"SIRE", "APAUII", "APAUBIO", "PIC", "EJIA"}

    materias = append(materias, "CCPD")

    for _, materia := range materias {
        fmt.Println(materia)
    }
}

```

## Ejercicio 7
Enunciado: Cuenta cuantas letras tiene una palabra.
Pistas: recorre la cadena con `range`.
Practica: strings, runes, bucles.


```go

package main

import(
    "fmt"
)

func main(){
    contador := 0
    palabra := hola

    for _, letra in range palabra{
        contador++;
    }

    fmt.Println(contador)
}

```

## Ejercicio 8
Enunciado: Usa un `map` para guardar la edad de tres personas.
Pistas: claves `string`, valores `int`.
Practica: maps.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 9
Enunciado: Consulta una clave que exista y otra que no exista en un `map`.
Pistas: revisa el segundo valor booleano.
Practica: acceso seguro a maps.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 10
Enunciado: Define un `struct` `Curso` con nombre, profesor y creditos.
Pistas: crea una instancia y muestra sus campos.
Practica: structs.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 11
Enunciado: Crea un metodo `Descripcion` para el struct `Curso`.
Pistas: usa un receptor de valor.
Practica: metodos.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 12
Enunciado: Escribe una funcion que reciba un puntero y multiplique su valor por dos.
Pistas: usa `*` y `&`.
Practica: punteros.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 13
Enunciado: Declara una constante para la tasa de descuento de una tienda.
Pistas: usa `const`.
Practica: constantes.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 14
Enunciado: Crea un programa que convierta grados Celsius a Fahrenheit.
Pistas: define una funcion.
Practica: funciones, tipos numericos.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 15
Enunciado: Imprime tres mensajes en orden con un `switch` simple.
Pistas: usa una variable numerica.
Practica: switch.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 16
Enunciado: Modela un `Libro` con autor y paginas, y muestra un resumen.
Pistas: struct y funcion auxiliar.
Practica: structs, funciones.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 17
Enunciado: Crea un array de 3 enteros y muestra su contenido.
Pistas: usa un tamano fijo.
Practica: arrays.


```go

package main

import(
    "fmt"
)

func main(){

}

```

## Ejercicio 18
Enunciado: Copia un slice y modifica la copia sin tocar el original.
Pistas: usa `copy`.
Practica: slices, copia de datos.


```go

package main

import(
    "fmt"
)

func main(){

}

```
