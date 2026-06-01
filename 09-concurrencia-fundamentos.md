# 09. Fundamentos de concurrencia

La concurrencia es la capacidad de gestionar varias tareas que progresan en el tiempo. No significa necesariamente ejecutarlas al mismo instante. El paralelismo, en cambio, si implica ejecucion simultanea en varios nucleos cuando el hardware lo permite. La asincronia describe una forma de no bloquear una tarea mientras otra sigue su curso.

## Diferencias clave

```text
Concurrencia:  A y B avanzan intercaladas.
Paralelismo:   A y B ejecutan al mismo tiempo en hardware distinto.
Asincronia:    lanzo una tarea y sigo haciendo otra cosa.
```

## Analogia intuitiva

Una cocina con un solo cocinero puede ser concurrente si organiza varias recetas en pasos. Una cocina con varios cocineros ademas puede ser paralela. Si una receta se deja en el horno y el cocinero sigue con otra, hay asincronia.

## Por que importa en Go

Go fue disenado para que crear tareas concurrentes sea barato y natural. En lugar de depender de hilos pesados y APIs verbosas, usa `goroutines`, `channels` y primitivas de sincronizacion simples.

## Ejemplo visual

```text
Tiempo ->
Tarea A: [trabaja]---[espera]---[trabaja]
Tarea B: ----[trabaja]---[espera]---[trabaja]
```

## Mini ejemplo

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Inicio")
    go fmt.Println("Trabajo concurrente")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Fin")
}
```

## Comparacion con Java

- Java tambien tiene concurrencia, pero Go la hace mas ligera y directa.
- En Java suele ser comun pensar en pools y futuros desde el principio.
- En Go la concurrencia suele empezar con `go` y luego se coordina con canales y sincronizacion.

## Errores frecuentes

- Confundir concurrencia con paralelismo.
- Pensar que crear una tarea concurrente ya resuelve la coordinacion.
- Olvidar que las tareas concurrentes necesitan control de vida y cierre.

## Buenas practicas

- Identifica primero la unidad de trabajo.
- Piensa despues en quien produce, quien consume y quien cancela.
- No conviertas cada problema secuencial en concurrente sin necesidad.

## Preguntas de autoevaluacion

1. Que diferencia real hay entre concurrencia y paralelismo?
2. En que casos la asincronia mejora la experiencia de usuario?
3. Por que la coordinacion es tan importante como la creacion de tareas?

## Ejercicio guiado

Describe con tus palabras un caso cotidiano donde exista concurrencia pero no paralelismo.

## Ejercicio sin resolver

Piensa en una aplicacion de pedidos y explica que partes serian concurrentes, cuales paralelas y cuales asincronas.
