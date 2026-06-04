# Ejercicios avanzados

## Ejercicio 33
Enunciado: Lanza una goroutine que imprima un mensaje y espera correctamente a que termine.
Pistas: usa `WaitGroup`.
Practica: goroutines, sincronizacion.

```go
package main

import (
    "fmt"
    "sync"
)

func imprimirMensaje(wg *sync.WaitGroup){
    defer wg.Done()
    fmt.Println("Mensaje")
}

func main(){
    // alternativa: var wg sync.waitGroup
    wg := sync.WaitGroup{}
    wg.Add(1)
    go imprimirMensaje(&wg)
    wg.Wait()
}
```

## Ejercicio 34
Enunciado: Crea dos goroutines que envien datos a un channel y recoge los resultados.
Pistas: piensa en productor-consumidor.
Practica: channels, concurrencia.

```go
package main

import (
    "fmt"
    "sync"
)

// escritura
func producir(ch chan<- int, wg *sync.WaitGroup){
    defer wg.Done()
    ch <- 1
}

// lectura
func consumir(ch <-chan int, wg *sync.WaitGroup){
    defer wg.Done()
    fmt.Println(<-ch)
}

func main(){
    wg := sync.WaitGroup{}
    ch := make(chan int)
    
    wg.Add(4)
    go producir(ch, &wg)
    go producir(ch, &wg)

    // wg.Wait() hace deadlock los productores se quedan bloqueados intentando enviar: ch <- 1 y nunca ejecutan wg.Done()

    go consumir(ch, &wg)
    go consumir(ch, &wg)
    
    wg.Wait()
}
```

## Ejercicio 35
Enunciado: Implementa un `select` con un timeout de un segundo.
Pistas: usa `time.After`.
Practica: select, timeout.

```go
package main

import (
    "fmt"
    "time"
    "sync"
)

func doWork(ch chan<- int, wg *sync.WaitGroup){
    defer wg.Done()
    time.Sleep(2 * time.Second)
    ch <- 1
}

func main(){
    ch := make(chan int)
    var wg sync.WaitGroup
    wg.Add(1)

    go doWork(ch, &wg)

    select{
        case entero := <-ch:
            fmt.Println(entero)
        case <-time.After(1 * time.Second):
            fmt.Println("timeout")
    }

    wg.Wait()

}
```

## Ejercicio 36
Enunciado: Crea un buffered channel y explica por que no bloquea en los primeros envios.
Pistas: prueba diferentes capacidades.
Practica: buffered channels.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 37
Enunciado: Implementa un worker pool con tres workers y cinco tareas.
Pistas: separa jobs y results.
Practica: patron worker pool.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 38
Enunciado: Protege un contador compartido con `Mutex`.
Pistas: revisa cada acceso al dato.
Practica: mutex, race conditions.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 39
Enunciado: Diseña una cache concurrente de solo lectura frecuente con `RWMutex`.
Pistas: distingue lector y escritor.
Practica: RWMutex.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 40
Enunciado: Usa `context.WithCancel` para detener un worker.
Pistas: escucha `ctx.Done()`.
Practica: context, cancelacion.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 41
Enunciado: Usa `context.WithTimeout` para abortar una tarea lenta.
Pistas: responde al timeout.
Practica: context, timeout.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 42
Enunciado: Construye un pipeline de dos etapas con channels.
Pistas: una etapa transforma y otra agrega.
Practica: pipeline.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 43
Enunciado: Implementa fan-out repartiendo tareas entre varias goroutines.
Pistas: comparte la entrada.
Practica: fan-out.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 44
Enunciado: Implementa fan-in unificando salidas de varias goroutines.
Pistas: usa un canal de merge.
Practica: fan-in.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 45
Enunciado: Detecta y corrige una race condition en un contador global.
Pistas: ejecuta mentalmente varias intercalaciones.
Practica: diagnostico de concurrencia.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 46
Enunciado: Evita una goroutine leak en un consumidor de channels.
Pistas: define cierre y cancelacion.
Practica: leaks, cierre correcto.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 47
Enunciado: Reescribe una solucion basada en estado compartido usando channels.
Pistas: compara claridad de ambas versiones.
Practica: diseno concurrente.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 48
Enunciado: Crea un sistema que procese tareas y pueda ser cancelado desde fuera.
Pistas: combina context y select.
Practica: cancelacion cooperativa.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 49
Enunciado: Simula una operacion remota con timeout y manejo de error.
Pistas: no bloquees indefinidamente.
Practica: context, select, errores.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 50
Enunciado: Añade estadisticas protegidas por mutex a un worker pool.
Pistas: separa datos compartidos y flujo de mensajes.
Practica: mutex, patrones.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 51
Enunciado: Implementa un lector de eventos con `select` que cierre correctamente al terminar.
Pistas: evita canales colgantes.
Practica: select, shutdown limpio.

```go
package main

import (
    "fmt"
)

func main(){

}
```

## Ejercicio 52
Enunciado: Diseña una solucion concurrente para calcular cuadrados de muchos numeros sin saturar recursos.
Pistas: worker pool o fan-out controlado.
Practica: escalado, limitacion.

```go
package main

import (
    "fmt"
)

func main(){

}
```
