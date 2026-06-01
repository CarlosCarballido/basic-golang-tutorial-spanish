# 16. Context

`context` permite propagar cancelacion, tiempos limite y valores de peticion a traves de varias goroutines.

## Ideas principales

- Cancelacion cooperativa.
- Timeout o deadline.
- Propagacion a funciones hijas.

## Analogia intuitiva

Un context es como una orden general de trabajo que acompana a toda una operacion. Si el supervisor cancela la orden, todas las tareas relacionadas deben parar cuanto antes.

## Ejemplo minimo

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

## Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

## Ejemplo practico

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func trabajo(ctx context.Context, salida chan<- string) {
    select {
    case <-time.After(3 * time.Second):
        salida <- "terminado"
    case <-ctx.Done():
        salida <- "cancelado"
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    salida := make(chan string)
    go trabajo(ctx, salida)

    fmt.Println(<-salida)
}
```

### Linea por linea

- `context.WithTimeout(...)`: crea un context con tiempo limite.
- `ctx.Done()`: canal que se cierra al cancelar o agotar el plazo.
- `select`: el worker responde al evento de tiempo o cancelacion.

## Comparacion con Java

- Se parece a pasar una senal de cancelacion por el flujo, pero en Go es un patron estandar.
- En Java la cancelacion suele depender mas de librerias o de la arquitectura concreta.

## Errores frecuentes

- Usar context para guardar datos arbitrarios sin necesidad.
- Olvidar llamar al `cancel()` devuelto por los contextos con timeout o cancelacion.
- No propagar el context a funciones hijas.

## Buenas practicas

- Pasa `context.Context` como primer parametro.
- No guardes context en structs.
- Cancela siempre que la operacion ya no tenga sentido.

## Preguntas de autoevaluacion

1. Por que `context` es central en APIs concurrentes?
2. Que diferencia hay entre cancelacion y timeout?
3. Cuando debes llamar a `cancel()`?

## Ejercicio guiado

Escribe una funcion que haga una tarea y se detenga si el context se cancela.

## Ejercicio sin resolver

Construye una operacion de tres pasos donde cada paso respete el mismo context.
