# 19. Ticker en Go

Un `ticker` es una utilidad del paquete `time` que genera eventos periodicos. Se usa cuando quieres ejecutar una accion cada cierto intervalo, por ejemplo para enviar latidos, refrescar una pantalla, medir progreso o disparar una tarea recurrente.

## Idea central

```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

for range ticker.C {
    fmt.Println("evento periodico")
}
```

`ticker.C` es un canal que entrega un valor cada vez que vence el intervalo.

## Analogia intuitiva

Piensa en un metrónomo. No hace una sola marca de tiempo y ya esta. Va marcando el pulso una y otra vez hasta que le dices que se detenga.

## Para que sirve

- Ejecutar una tarea cada N segundos.
- Informar progreso periodico.
- Enviar heartbeats en sistemas distribuidos.
- Hacer sondeos regulares sin bucles manuales con `Sleep`.

## Ejemplo minimo

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(500 * time.Millisecond)
    defer ticker.Stop()

    for i := 0; i < 3; i++ {
        <-ticker.C
        fmt.Println("pulso", i+1)
    }
}
```

### Explicacion linea por linea

- `time.NewTicker(500 * time.Millisecond)`: crea un ticker que marca cada medio segundo.
- `defer ticker.Stop()`: libera recursos cuando termina el programa.
- `<-ticker.C`: espera al siguiente pulso del ticker.
- `for i := 0; i < 3; i++`: limita el numero de iteraciones para que el programa finalice.

## Ejemplo realista

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    progreso := 0
    for range ticker.C {
        progreso += 20
        fmt.Printf("Progreso: %d%%\n", progreso)
        if progreso >= 100 {
            break
        }
    }
}
```

## Diferencia entre `Ticker` y `Timer`

- `Timer` dispara una sola vez.
- `Ticker` dispara repetidamente cada intervalo.

## Diferencia entre `Ticker` y `time.Sleep`

- `time.Sleep` pausa la goroutine durante un tiempo fijo.
- `Ticker` genera una secuencia de señales periodicas.
- `Ticker` es mejor cuando necesitas un ritmo repetido y no solo una espera unica.

## Cuando usarlo

- Tareas repetitivas con intervalo estable.
- Sistemas que necesitan latidos regulares.
- Lectura periodica de colas, metricas o estado.

## Cuando no usarlo

- Si solo necesitas esperar una vez, usa `Timer` o `Sleep`.
- Si el intervalo depende de la respuesta de otro sistema, puede que necesites `select`, no un ticker fijo.
- Si el trabajo tarda mas que el intervalo, debes decidir si acumulas ticks o si ignoras los atrasos.

## Ejemplo combinado con `select`

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    stop := time.After(3500 * time.Millisecond)

    for {
        select {
        case <-ticker.C:
            fmt.Println("haciendo trabajo periodico")
        case <-stop:
            fmt.Println("fin")
            return
        }
    }
}
```

## Comparacion con Java

- En Java esto suele resolverse con `ScheduledExecutorService` o `TimerTask`.
- En Go el `Ticker` es directo y encaja con la filosofia de canales y `select`.

## Errores frecuentes

- Olvidar llamar a `Stop()` y dejar recursos activos.
- Usar `Ticker` para una accion unica.
- Bloquearse esperando ticks sin una condicion de salida.
- Creer que un ticker compensa un diseno sin cancelacion.

## Buenas practicas

- Siempre detiene el ticker con `defer ticker.Stop()`.
- Combinalo con `select` para permitir cancelacion o timeout.
- Elige intervalos realistas para el trabajo que vas a hacer.
- Si el trabajo puede tardar mas que el intervalo, piensa en la estrategia de acumulacion o descarte.

## Preguntas de autoevaluacion

1. Que diferencia hay entre un `Ticker` y un `Timer`?
2. Por que conviene llamar a `Stop()`?
3. En que casos es mejor `Ticker` que un bucle con `Sleep`?

## Ejercicio guiado

Escribe un programa que imprima un mensaje cada segundo durante cinco ticks y luego termine.

## Ejercicio sin resolver

Construye un monitor de estado que use `Ticker` para imprimir una estadistica periodica y `select` para detenerse con `context`.
