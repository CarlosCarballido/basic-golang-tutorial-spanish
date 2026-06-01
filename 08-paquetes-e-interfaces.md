# 08. Paquetes e interfaces

Los paquetes organizan el codigo en unidades coherentes. Las interfaces definen comportamiento por capacidad, no por linaje. Esto es una de las ideas mas importantes de Go.

## Paquetes

Cada archivo pertenece a un paquete. Un paquete puede contener funciones, tipos y constantes relacionadas.

```go
package utilidades
```

## Exportacion

En Go, un identificador que empieza por mayuscula se exporta. Uno que empieza por minuscula queda privado al paquete.

## Interfaces

```go
type Imprimible interface {
    Imprimir() string
}
```

Un tipo implementa una interfaz simplemente si tiene los metodos requeridos. No necesita declararlo explicitamente.

## Analogia intuitiva

Una interfaz es como una toma universal: si el enchufe encaja y entrega la funcion esperada, sirve. No importa si el aparato fue fabricado en otra linea.

## Ejemplo completo

```go
package main

import "fmt"

type Figura interface {
    Area() float64
}

type Cuadrado struct {
    Lado float64
}

func (c Cuadrado) Area() float64 {
    return c.Lado * c.Lado
}

func mostrarArea(f Figura) {
    fmt.Println(f.Area())
}

func main() {
    mostrarArea(Cuadrado{Lado: 4})
}
```

### Linea por linea

- `type Figura interface`: define el contrato.
- `Area() float64`: metodo requerido.
- `mostrarArea(f Figura)`: acepta cualquier tipo que cumpla el contrato.

## Comparacion con Java

- En Java la clase debe declarar `implements`.
- En Go la implementacion es implicita.
- Go favorece interfaces pequenas y centradas en comportamiento concreto.

## Errores frecuentes

- Crear interfaces demasiado grandes.
- Pensar en interfaces antes de necesitar polimorfismo.
- Confundir paquete con modulo.

## Buenas practicas

- Define interfaces donde se consumen, no donde se implementan.
- Mantiene interfaces pequenas.
- Usa nombres que expresen comportamiento, no tecnologia.

## Preguntas de autoevaluacion

1. Por que las interfaces pequenas son mejores en Go?
2. Que significa que una implementacion sea implicita?
3. Como ayuda esto a probar codigo?

## Ejercicio guiado

Define una interfaz `Notificador` con un metodo `Enviar` y crea una implementacion simple por consola.

## Ejercicio sin resolver

Disena un paquete `dominio` con un `struct` y una interfaz que permita procesar pagos.
