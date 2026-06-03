# Ejercicios intermedios

## Ejercicio 19
Enunciado: Define una interfaz `Figure` con un metodo `Area`.
Pistas: crea dos implementaciones distintas.
Practica: interfaces, polimorfismo.

```go
package main

import(
    "fmt"
)

type Figure interface{
    Area() float64
}

type Rectangulo struct{
    x float64
    y float64
}

func (r Rectangulo) Area() float64 {
    return r.x * r.y
}

type Circulo struct {
	radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.radio * c.radio
}

func mostrarArea(f Figure) {
	fmt.Println("Área:", f.Area())
}

func main() {
    r := Rectangulo{3, 5}
	c := Circulo{2}

	mostrarArea(r)
	mostrarArea(c)
}
```

## Ejercicio 20
Enunciado: Construye un paquete utilitario con una funcion exportada y otra privada.
Pistas: controla mayusculas y minusculas.
Practica: paquetes, exportacion.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 21
Enunciado: Crea un metodo con receptor por puntero que modifique un campo de un struct.
Pistas: compara antes y despues.
Practica: metodos, punteros.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 22
Enunciado: Escribe una funcion que devuelva dos valores: resultado y error.
Pistas: simula una division segura.
Practica: retorno multiple, errores.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 23
Enunciado: Modela una `CuentaBancaria` con deposito y retiro.
Pistas: usa receptor por puntero.
Practica: structs, metodos.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 24
Enunciado: Crea una funcion que reciba un slice y devuelva la suma total.
Pistas: recorre con `range`.
Practica: slices, funciones.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 25
Enunciado: Crea una funcion que reciba un mapa de notas y devuelva la nota media.
Pistas: recorre claves y valores.
Practica: maps, funciones.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 26
Enunciado: Diseña un tipo `Rectangulo` y dos metodos: `Area` y `Perimetro`.
Pistas: evita calculos duplicados.
Practica: structs, metodos.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 27
Enunciado: Crea una interfaz `Logger` con metodo `Log` y una implementacion por consola.
Pistas: piensa en comportamiento, no en tecnologia.
Practica: interfaces.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 28
Enunciado: Crea un paquete `mathx` con una funcion `Max`.
Pistas: prueba a importarlo desde main.
Practica: paquetes.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 29
Enunciado: Escribe una funcion que reciba un puntero a `int` y lo ponga a cero.
Pistas: evita copiar el valor.
Practica: punteros.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 30
Enunciado: Crea un struct `Alumno` que embeba `Persona`.
Pistas: accede a campos heredados por composicion.
Practica: composicion.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 31
Enunciado: Implementa una funcion que reciba una lista de strings y devuelva las no vacias.
Pistas: crea un nuevo slice.
Practica: slices, filtrado.

```go
package main

import(
    "fmt"
)

func main{

}
```

## Ejercicio 32
Enunciado: Crea una funcion que lea una configuracion desde un `map[string]string`.
Pistas: comprueba si la clave existe.
Practica: maps, robustez.

```go
package main

import(
    "fmt"
)

func main{

}
```
