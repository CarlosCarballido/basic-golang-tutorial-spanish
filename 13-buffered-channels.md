# 13. Buffered channels

Un channel buffered tiene capacidad interna. Puede almacenar varios valores antes de bloquear al emisor.

## Sintaxis

```go
ch := make(chan int, 3)
```

## Que cambia

- Sin buffer: el envio espera a que exista receptor.
- Con buffer: el envio avanza hasta llenar la capacidad.

## Analogia intuitiva

Piensa en una caja con espacio para tres paquetes. Mientras haya sitio, el repartidor sigue dejando paquetes. Cuando la caja se llena, debe esperar.

## Ejemplo minimo

```go
ch := make(chan string, 2)
ch <- "a"
ch <- "b"
fmt.Println(<-ch)
fmt.Println(<-ch)
```

## Cuándo utilizarlo

- Para desacoplar ligeramente productor y consumidor.
- Para absorber picos breves de carga.
- Para limitar una cola interna con capacidad conocida.

## Cuándo NO utilizarlo

- Cuando el orden de coordinacion debe ser estricto y simple.
- Cuando el buffer solo oculta un problema de diseno.
- Cuando no sabes por que elegiste esa capacidad.

## Ejemplo realista

```go
package main

import "fmt"

func main() {
    jobs := make(chan int, 2)

    jobs <- 1
    jobs <- 2

    fmt.Println(<-jobs)
    fmt.Println(<-jobs)
}
```

## Comparacion con Java

- Se parece conceptualmente a una cola con capacidad maxima.
- La diferencia es que aqui el canal tambien forma parte del mecanismo de sincronizacion.

## Errores frecuentes

- Confundir buffer grande con mejor diseno.
- Olvidar que un buffer finito tambien se llena.
- Creer que un buffered channel elimina por completo el bloqueo.

## Buenas practicas

- Elige capacidad por razon de negocio, no por intuicion vaga.
- Documenta por que existe el buffer.
- Usa el buffer como herramienta de ritmo, no como parche permanente.

## Preguntas de autoevaluacion

1. Que problema resuelve un buffered channel?
2. Que ocurre cuando el buffer se llena?
3. Por que un buffer excesivo puede ser mala idea?

## Ejercicio guiado

Crea un channel buffered de capacidad 3, envia tres elementos y comprueba que el cuarto envio bloquea hasta que se lea uno.

## Ejercicio sin resolver

Disena un pequeño sistema de cola con un buffered channel y un consumidor lento.
