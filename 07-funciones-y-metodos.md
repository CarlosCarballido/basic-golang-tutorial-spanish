# 07. Funciones y metodos

Las funciones son bloques reutilizables de codigo. Los metodos son funciones asociadas a un tipo concreto. En Go esto ayuda a modelar comportamiento sin convertir todo en una jerarquia de clases.

## Funcion basica

```go
func suma(a int, b int) int {
    return a + b
}
```

## Parametros y retorno multiple

```go
func dividir(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division por cero")
    }
    return a / b, nil
}
```

## Analogia intuitiva

Una funcion es una maquina con entradas y salidas. Un metodo es esa misma maquina, pero montada sobre una pieza concreta del sistema.

## Metodo

```go
type Rectangulo struct {
    Ancho, Alto float64
}

func (r Rectangulo) Area() float64 {
    return r.Ancho * r.Alto
}
```

## Ejemplo realista

```go
package main

import "fmt"

type Cuenta struct {
    Titular string
    Saldo   float64
}

func (c *Cuenta) Depositar(monto float64) {
    c.Saldo += monto
}

func (c Cuenta) Resumen() string {
    return fmt.Sprintf("%s: %.2f", c.Titular, c.Saldo)
}

func main() {
    cuenta := Cuenta{Titular: "Sara", Saldo: 100}
    cuenta.Depositar(50)
    fmt.Println(cuenta.Resumen())
}
```

### Linea por linea

- `func (c *Cuenta) Depositar`: metodo con receptor por puntero para mutar saldo.
- `func (c Cuenta) Resumen()`: metodo con receptor por valor para solo leer.
- `cuenta.Depositar(50)`: actualiza el objeto original.

## Comparacion con Java

- Java usa clases y metodos como unidad central.
- Go separa datos y comportamiento de forma mas flexible.
- Un metodo en Go puede usar receptor por valor o por puntero, segun convenga.

## Errores frecuentes

- Usar receptor por valor cuando necesitas modificar estado.
- Hacer funciones demasiado largas.
- No devolver errores cuando una operacion puede fallar.

## Buenas practicas

- Mantiene funciones pequenas y centradas.
- Devuelve errores en vez de ocultar fallos.
- Usa receptor por puntero si el metodo modifica el estado.

## Preguntas de autoevaluacion

1. Cuando un metodo deberia usar receptor por puntero?
2. Por que Go favorece el retorno multiple?
3. Que diferencia hay entre funcion y metodo?

## Ejercicio guiado

Escribe una funcion `promedio` que reciba un slice de enteros y devuelva un `float64`.

## Ejercicio sin resolver

Crea un metodo `Descuento` sobre un `Producto` que reduzca su precio en un porcentaje.
