# 10. Goroutines

Una goroutine es una funcion que se ejecuta concurrentemente con otras goroutines. Se crea con la palabra clave `go`.

## Ciclo de vida

```text
crear -> ejecutar -> bloquearse -> reanudar -> terminar
```

## Analogia intuitiva

Una goroutine es como abrir una nueva tarea en una mesa de trabajo. No es un proceso pesado; es una unidad ligera que el runtime puede mover y coordinar de manera eficiente.

## Ejemplo minimo

```go
package main

import (
    "fmt"
    "time"
)

func saludar() {
    fmt.Println("Hola desde una goroutine")
}

func main() {
    go saludar()
    time.Sleep(50 * time.Millisecond)
}
```

## Multiples goroutines

```go
for i := 1; i <= 3; i++ {
    go func(n int) {
        fmt.Println("Tarea", n)
    }(i)
}
```

## Miles de goroutines

Go puede gestionar miles de goroutines porque el coste es mucho menor que el de muchos hilos del sistema operativo. Aun asi, "barato" no significa "gratis": demasiadas tareas sin control pueden saturar CPU, memoria o canales.

## Scheduler de Go

El runtime reparte goroutines sobre hilos del sistema operativo. El programador normalmente no decide ese reparto de forma manual. Esa abstraccion es una de las grandes ventajas del lenguaje.

## Ejemplo realista

```go
package main

import (
    "fmt"
    "sync"
)

func trabajo(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Procesando", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go trabajo(i, &wg)
    }
    wg.Wait()
}
```

## Comparacion con Java

- Una goroutine se parece conceptualmente a una tarea ligera, no a un hilo explicito pesado.
- En Java el programador suele pensar antes en `Thread`, `ExecutorService` o `CompletableFuture`.
- En Go la abstraccion base es mucho mas pequena.

## Errores frecuentes

- Lanzar goroutines sin una forma clara de esperar su finalizacion.
- Capturar variables de bucle de forma incorrecta.
- Crear goroutines que nunca terminan.

## Buenas practicas

- Usa `WaitGroup` o `context` para coordinar finalizacion.
- Pasa variables como argumento al cierre cuando vengas de un bucle.
- Limita el numero de goroutines activas cuando el problema lo requiera.

## Preguntas de autoevaluacion

1. Por que una goroutine es barata comparada con un hilo tradicional?
2. Que problema evita pasar `i` como argumento al closure?
3. Cuando necesitas esperar a una goroutine?

## Ejercicio guiado

Lanza tres goroutines que impriman mensajes distintos y espera a que todas terminen.

## Ejercicio sin resolver

Crea una funcion que lance 100 goroutines y contabiliza cuantas terminan correctamente.
