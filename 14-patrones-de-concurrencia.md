# 14. Patrones de concurrencia

Los patrones de concurrencia son soluciones recurrentes para organizar goroutines, channels y cancelacion. No son recetas rigidas; son estructuras que ayudan a pensar en sistemas reales.

## Worker Pool

```text
jobs -> [worker1]
jobs -> [worker2]
jobs -> [worker3]
```

### Idea

Un conjunto fijo de workers procesa una cola de tareas. Sirve para limitar consumo de recursos y controlar paralelismo.

### Caso de uso

Procesamiento de imagenes, validacion masiva, peticiones a servicios externos con limite de concurrencia.

## Fan-Out

```text
        -> worker A
input -> -> worker B
        -> worker C
```

### Idea

Una entrada se reparte entre varios consumidores.

### Caso de uso

Distribuir calculos independientes.

## Fan-In

```text
worker A ->
worker B -> merge -> salida
worker C ->
```

### Idea

Varias salidas se combinan en una sola salida.

### Caso de uso

Recolectar resultados de multiples workers.

## Pipeline

```text
fuente -> etapa1 -> etapa2 -> etapa3 -> resultado
```

### Idea

Cada etapa transforma datos y pasa el resultado a la siguiente.

### Caso de uso

Extraccion, limpieza y analisis de datos.

## Producer Consumer

```text
producer -> channel -> consumer
```

### Idea

El productor genera trabajo y el consumidor lo procesa.

### Caso de uso

Colas de eventos, logs, notificaciones.

## Cancellation Pattern

```text
context cancelado -> todos los workers paran
```

### Idea

Una senal comun detiene el trabajo pendiente.

### Caso de uso

Cancelar una operacion cuando falla una etapa critica.

## Timeout Pattern

```text
si no llega respuesta a tiempo -> abandonar
```

### Idea

El sistema no espera indefinidamente.

### Caso de uso

Llamadas de red, esperas a servicios remotos, UI reactiva.

## Ejemplo completo: worker pool

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int)
    results := make(chan int)
    var wg sync.WaitGroup

    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    go func() {
        for j := 1; j <= 5; j++ {
            jobs <- j
        }
        close(jobs)
    }()

    go func() {
        wg.Wait()
        close(results)
    }()

    for r := range results {
        fmt.Println(r)
    }
}
```

## Buenas practicas generales

- Elige el patron mas simple que resuelva el problema.
- Dibuja primero el flujo de datos.
- Asegura el cierre correcto de canales y goroutines.

## Preguntas de autoevaluacion

1. Cuando un worker pool es mejor que lanzar una goroutine por tarea?
2. Que diferencia hay entre fan-out y fan-in?
3. Por que los pipelines facilitan el razonamiento?

## Ejercicio guiado

Identifica que patron usarias para procesar 10.000 registros sin saturar la memoria.

## Ejercicio sin resolver

Disena un pipeline de tres etapas para procesar texto de entrada, normalizarlo y producir estadisticas.
