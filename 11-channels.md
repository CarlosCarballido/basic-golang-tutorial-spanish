# 11. Channels

Los channels son conductos tipados para comunicar goroutines. Sirven para enviar y recibir datos y, ademas, para coordinar el ritmo entre productores y consumidores.

## Sintaxis basica

```go
ch := make(chan int)
ch <- 10
valor := <-ch
```

## Diagrama ASCII

```text
Main ---> channel ---> Worker
```

## Analogia intuitiva

Un channel es como una cinta transportadora entre dos puestos de trabajo. El productor deja un paquete y el consumidor lo recoge. Si la cinta esta vacia o llena, uno de los lados espera segun el tipo de canal.

## Ejemplo minimo

```go
package main

import "fmt"

func main() {
    ch := make(chan string)
    go func() { ch <- "mensaje" }()
    fmt.Println(<-ch)
}
```

### Linea por linea

- `make(chan string)`: crea un channel sin buffer.
- `go func() { ch <- "mensaje" }()`: una goroutine envia un valor.
- `fmt.Println(<-ch)`: la goroutine principal recibe el valor.

## Cierre de channels

```go
close(ch)
for v := range ch {
    fmt.Println(v)
}
```

Cerrar un channel comunica que no llegaran mas valores. Solo quien produce deberia cerrarlo.

## Ejemplo realista

```go
package main

import "fmt"

func worker(entrada <-chan int, salida chan<- int) {
    for n := range entrada {
        salida <- n * n
    }
    close(salida)
}

func main() {
    entrada := make(chan int)
    salida := make(chan int)

    go worker(entrada, salida)

    go func() {
        defer close(entrada)
        for i := 1; i <= 3; i++ {
            entrada <- i
        }
    }()

    for v := range salida {
        fmt.Println(v)
    }
}
```

## Comparacion con Java

- En Java suele usarse una cola concurrente o un executor para un efecto parecido.
- En Go el channel forma parte del lenguaje y comunica tanto datos como sincronizacion.
- Esto simplifica mucho los flujos producer-consumer.

## Errores frecuentes

- Enviar en un channel cerrado.
- Cerrar un channel desde el lado incorrecto.
- Olvidar que un channel sin buffer bloquea al emisor hasta que haya receptor.

## Buenas practicas

- Usa channels para comunicar, no para compartir estado cuando no hace falta.
- Documenta claramente quien produce, quien consume y quien cierra.
- Usa channels de recepcion o envio cuando quieras expresar intencion.

## Preguntas de autoevaluacion

1. Que diferencia hay entre un channel con y sin buffer?
2. Quien deberia cerrar un channel?
3. Que sucede si nadie recibe de un channel sin buffer?

## Ejercicio guiado

Crea un channel de enteros, envia tres valores y recibilos en la goroutine principal.

## Ejercicio sin resolver

Construye un pipeline de dos etapas usando channels para transformar texto a mayusculas y luego contar caracteres.
