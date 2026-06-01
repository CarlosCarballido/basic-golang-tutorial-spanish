# 12. Select

`select` permite esperar sobre varios channels al mismo tiempo. Es la pieza que hace posible la multiplexacion de eventos, timeouts y cancelacion cooperativa.

## Idea central

```go
select {
case v := <-ch1:
    fmt.Println(v)
case ch2 <- 10:
    fmt.Println("enviado")
default:
    fmt.Println("sin actividad")
}
```

## Analogia intuitiva

Es como esperar en varias ventanillas a la vez. En cuanto una tiene disponibilidad, se atiende esa opcion. Si ninguna esta lista y existe `default`, el programa no se bloquea.

## Ejemplo con timeout

```go
select {
case resultado := <-trabajo:
    fmt.Println("Resultado:", resultado)
case <-time.After(2 * time.Second):
    fmt.Println("Tiempo agotado")
}
```

## Ejemplo con default

```go
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("No hay mensajes")
}
```

## Ejemplo realista

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    rapido := make(chan string)
    lento := make(chan string)

    go func() { time.Sleep(200 * time.Millisecond); rapido <- "rapido" }()
    go func() { time.Sleep(500 * time.Millisecond); lento <- "lento" }()

    for i := 0; i < 2; i++ {
        select {
        case v := <-rapido:
            fmt.Println(v)
        case v := <-lento:
            fmt.Println(v)
        }
    }
}
```

## Comparacion con Java

- En Java la multiplexacion suele resolverse con `CompletableFuture`, `Selector` o librerias adicionales.
- En Go `select` es nativo y encaja directamente con channels.

## Errores frecuentes

- Usar `default` cuando si necesitas bloquear hasta recibir datos.
- No considerar que `select` elige aleatoriamente entre casos listos.
- Olvidar que `time.After` crea recursos temporales.

## Buenas practicas

- Usa `select` para coordinar entrada de datos, cancelacion y timeouts.
- Si la espera debe ser obligatoria, no metas `default` por costumbre.
- Piensa en la semantica del timeout antes de codificarlo.

## Preguntas de autoevaluacion

1. Que hace `default` en un `select`?
2. Por que `select` es util para timeouts?
3. Que pasa si varios casos estan listos?

## Ejercicio guiado

Escribe un `select` que imprima un mensaje si llega informacion en menos de un segundo y otro mensaje si no llega.

## Ejercicio sin resolver

Implementa un `select` que escuche dos canales de eventos y un canal de cancelacion.
