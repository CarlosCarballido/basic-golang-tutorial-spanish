# 20. Panic y recover en Go

`panic` y `recover` forman parte del mecanismo de manejo de fallos excepcionales en Go. No se usan para el control normal del programa, sino para situaciones realmente anormales: indices fuera de rango, acceso a un puntero `nil`, configuraciones imposibles o errores internos que no deberian continuar como si nada.

## Idea central

- `panic` detiene el flujo normal y empieza a deshacer la pila de llamadas.
- `recover` permite interceptar ese estado de panic dentro de una funcion diferida (`defer`).

## Analogia intuitiva

Piensa en una alarma de incendio. `panic` activa la alarma y obliga a evacuar. `recover` seria el equipo de emergencia que, en un punto controlado, puede parar la evacuacion y decidir si es seguro continuar.

## Cuándo usar `panic`

- En errores de programacion que no deberían ocurrir en ejecucion normal.
- En inicializacion critica cuando el programa no puede seguir sin una configuracion valida.
- En librerias internas cuando una invarianta fundamental se rompe.

## Cuándo NO usar `panic`

- Para errores esperables de negocio.
- Para validacion de formulario o entrada de usuario.
- Para reemplazar el retorno de `error`.

## Ejemplo minimo de `panic`

```go
package main

func main() {
    panic("algo salio muy mal")
}
```

### Explicacion linea por linea

- `panic("algo salio muy mal")`: interrumpe el flujo normal y muestra el error.
- El programa termina si nadie lo recupera.

## Ejemplo clasico de `recover`

```go
package main

import "fmt"

func seguro() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recuperado de:", r)
        }
    }()

    panic("fallo controlado")
}

func main() {
    seguro()
    fmt.Println("el programa continuo")
}
```

### Explicacion linea por linea

- `defer func() { ... }()`: registra una funcion que se ejecuta al salir de `seguro`.
- `recover()`: captura el valor del panic si la llamada ocurre dentro de un `defer`.
- `if r := recover(); r != nil`: comprueba si realmente hubo panic.
- `fmt.Println("el programa continuo")`: se ejecuta porque el panic fue manejado.

## Diagrama ASCII

```text
funcion normal
   |
   v
panic -> deshace pila -> ejecuta defer -> recover opcional -> salida controlada
```

## Ejemplo realista

```go
package main

import "fmt"

func dividir(a, b int) int {
    if b == 0 {
        panic("division por cero")
    }
    return a / b
}

func ejecutarOperacion() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("se detecto un fallo interno:", r)
        }
    }()

    fmt.Println(dividir(10, 0))
    fmt.Println("esto no se imprime si ocurre panic sin recover")
}

func main() {
    ejecutarOperacion()
}
```

## Diferencia con `error`

- `error` representa un fallo esperado y manejable.
- `panic` representa una situacion excepcional o una violacion grave de una invarianta.
- En Go se prefiere `error` para casi todo lo que el usuario pueda provocar.

## Comparacion con Java

- `panic` se parece solo parcialmente a lanzar una excepcion no capturada.
- `recover` no es igual a un `catch` general de uso cotidiano; es mas limitado y se usa dentro de `defer`.
- Go favorece `error` y reserva este mecanismo para casos excepcionales.

## Errores frecuentes

- Usar `panic` para validar datos de entrada.
- Intentar llamar a `recover` fuera de un `defer` y esperar que funcione.
- Recuperar un panic y ocultar el problema sin registrarlo.
- Pensar que `recover` evita todos los fallos posibles.

## Buenas practicas

- Usa `panic` con mucha moderacion.
- Recupera solo en los bordes del sistema o en una capa controlada.
- Loguea el motivo del panic antes de seguir.
- Si puedes devolver `error`, normalmente es mejor que hacer `panic`.

## Preguntas de autoevaluacion

1. Por que `recover` solo funciona dentro de un `defer`?
2. En que casos `panic` puede tener sentido?
3. Por que `error` sigue siendo la opcion preferida en Go?

## Ejercicio guiado

Escribe una funcion que haga `panic` si recibe un divisor cero y otra funcion que la llame con `recover` para evitar que el programa termine.

## Ejercicio sin resolver

Construye una funcion `safeExecute` que reciba una funcion como parametro, ejecute su logica y recupere cualquier `panic` para devolver un mensaje controlado.
