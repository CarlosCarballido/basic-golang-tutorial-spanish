# 18. Proyecto practico: Procesador concurrente de tareas

Este proyecto integra goroutines, channels, select, WaitGroup, Mutex y context en una aplicacion pequena pero realista.

## Objetivo

Construir un procesador concurrente de tareas que reciba trabajos, los distribuya entre workers, registre resultados, permita cancelar la ejecucion y respete un timeout global.

## Diseno

```text
entrada de tareas -> channel de jobs -> workers -> channel de resultados -> agregador
                         ^                                              |
                         |                                              v
                     context / cancelacion ----------------------> cierre limpio
```

### Componentes

- Productor de tareas.
- Pool de workers.
- Agregador de resultados.
- Contexto de cancelacion.
- Mutex para estadisticas compartidas.

## Implementacion

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

type Task struct {
    ID int
}

type Result struct {
    TaskID int
    Value  int
}

type Stats struct {
    mu       sync.Mutex
    Completed int
    Failed    int
}

func (s *Stats) IncCompleted() {
    s.mu.Lock()
    s.Completed++
    s.mu.Unlock()
}

func (s *Stats) IncFailed() {
    s.mu.Lock()
    s.Failed++
    s.mu.Unlock()
}

func worker(ctx context.Context, id int, tasks <-chan Task, results chan<- Result, stats *Stats, wg *sync.WaitGroup) {
    defer wg.Done()
    for {
        select {
        case <-ctx.Done():
            return
        case task, ok := <-tasks:
            if !ok {
                return
            }
            time.Sleep(100 * time.Millisecond)
            results <- Result{TaskID: task.ID, Value: task.ID * task.ID}
            stats.IncCompleted()
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    tasks := make(chan Task)
    results := make(chan Result)
    var stats Stats
    var wg sync.WaitGroup

    for i := 1; i <= 4; i++ {
        wg.Add(1)
        go worker(ctx, i, tasks, results, &stats, &wg)
    }

    go func() {
        defer close(tasks)
        for i := 1; i <= 10; i++ {
            select {
            case <-ctx.Done():
                stats.IncFailed()
                return
            case tasks <- Task{ID: i}:
            }
        }
    }()

    go func() {
        wg.Wait()
        close(results)
    }()

    for result := range results {
        fmt.Printf("task %d -> %d\n", result.TaskID, result.Value)
    }

    fmt.Printf("completed=%d failed=%d\n", stats.Completed, stats.Failed)
}
```

## Explicacion paso a paso

1. Se crea un `context` con timeout para que el sistema no espere indefinidamente.
2. Se crean dos channels: uno para tareas y otro para resultados.
3. Se lanzan workers en paralelo.
4. El productor envía tareas y respeta la cancelacion.
5. El agregador consume resultados hasta que el canal se cierra.
6. Un `Mutex` protege las estadisticas compartidas.
7. `WaitGroup` garantiza que los workers terminen antes de cerrar el canal de resultados.

## Mejoras posibles

- Limitar la capacidad de `tasks` con buffered channel.
- Separar errores de resultados exitosos.
- Añadir reintentos con backoff.
- Registrar tiempos por tarea.
- Exponer el procesador como paquete reutilizable.

## Ejercicios adicionales

1. Cambia el worker para que falle algunas tareas de forma aleatoria y contabiliza los fallos.
2. Sustituye el timeout global por un timeout por tarea.
3. Añade un canal de cancelacion manual para detener todo el sistema desde una condicion externa.
4. Convierte el programa en una libreria con una funcion `Run(ctx, tasks)`.

## Autoevaluacion final

1. Que pieza del proyecto coordina la cancelacion?
2. Por que se usa `WaitGroup` y no un canal extra para todo?
3. Donde aporta valor el `Mutex` y donde no?
