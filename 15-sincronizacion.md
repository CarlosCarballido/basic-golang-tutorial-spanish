# 15. Sincronizacion

La sincronizacion evita carreras de datos, bloqueos y estados inconsistentes cuando varias goroutines comparten recursos.

## WaitGroup

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
}()
wg.Wait()
```

### Idea

Permite esperar a que un grupo de goroutines termine.

### Comparacion con Java

Se parece conceptualmente a un `join()`, pero a escala de grupo, no de hilo individual.

## Mutex

```go
var mu sync.Mutex
mu.Lock()
contador++
mu.Unlock()
```

### Idea

Solo una goroutine accede a la seccion critica al mismo tiempo.

### Error comun

Olvidar hacer `Unlock()` incluso si algo falla.

## Ejemplo erroneo y corregido

```go
// Erroneo: acceso concurrente sin proteccion.
// contador++

// Correcto:
mu.Lock()
contador++
mu.Unlock()
```

## RWMutex

```go
var rw sync.RWMutex
rw.RLock()
valor := cache["clave"]
rw.RUnlock()
```

### Idea

Permite multiples lectores simultaneos y un solo escritor exclusivo.

### Cuándo mejora el rendimiento

Cuando hay muchas lecturas y pocas escrituras.

## Ejemplo realista

```go
package main

import (
    "fmt"
    "sync"
)

type Contador struct {
    mu sync.Mutex
    n  int
}

func (c *Contador) Inc() {
    c.mu.Lock()
    c.n++
    c.mu.Unlock()
}

func (c *Contador) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.n
}

func main() {
    var c Contador
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Inc()
        }()
    }
    wg.Wait()
    fmt.Println(c.Value())
}
```

## Comparacion con Java

- `Mutex` se parece a `synchronized` o a un lock explicito.
- `RWMutex` se parece a un lock de lectura/escritura.
- Go obliga a pensar con mas precision en la seccion critica.

## Errores frecuentes

- Proteger demasiado poco o demasiado.
- Dejar un lock adquirido si aparece un retorno temprano.
- Copiar un tipo que contiene un mutex.

## Buenas practicas

- Protege solo el estado compartido realmente necesario.
- Usa `defer` para liberar locks cuando la funcion es pequena.
- Considera channels si el problema se expresa mejor como comunicacion que como memoria compartida.

## Preguntas de autoevaluacion

1. Que problema resuelve un `Mutex`?
2. Cuando conviene usar `RWMutex`?
3. Por que copiar un mutex es peligroso?

## Ejercicio guiado

Convierte un contador inseguro en uno protegido con mutex.

## Ejercicio sin resolver

Implementa una cache concurrente con `RWMutex` y acceso seguro a lectura y escritura.
