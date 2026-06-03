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
package utilidades

import(
    "fmt"
)

func hola() {
    fmt.Println("Hola")
}

func Hola(){
    hola()
}
```

```go
package main

import(
    "fmt"
    "utilidades"
)

func main(){
    utilidades.Hola()
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

type Perro struct {
    patas int
}

func (p *Perro) modificarPatas(valor int){
    p.patas = valor
}

func main(){
    toby := Perro{4}

    fmt.Println("Antes:", toby.patas)
    toby.modificarPatas(3)
    fmt.Println("Después:", toby.patas)
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

func divisionSegura(a float64, b float64) (float64, error){
    if *b != 0{
        return *a / *b, nil
    } else {
        fmt.Println("No se puede dividir por 0")
        return 0, false
    }
}

func main(){
    a := 10.0
    b := 0.0

    divisionSegura(&a, &b)
}
```

## Ejercicio 23
Enunciado: Modela una `CuentaBancaria` con deposito y retiro.
Pistas: usa receptor por puntero.
Practica: structs, metodos.

```go
package main

import (
	"fmt"
)

type CuentaBancaria struct {
	dinero float64
}

func (c *CuentaBancaria) Deposito(cantidad float64) {
	c.dinero += cantidad
}

func (c *CuentaBancaria) Retiro(cantidad float64) {
	c.dinero -= cantidad
}

func main() {
	cuenta := CuentaBancaria{dinero: 100.0}

	fmt.Println("Saldo inicial:", cuenta.dinero)

	cuenta.Deposito(33.0)
	cuenta.Retiro(3.0)

	fmt.Println("Saldo final:", cuenta.dinero)
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

func sumReduce(slice []float64) float64{
    sumaTotal := 0.0
    for i:=0; i < len(slice) ; i++{
        sumaTotal += (slice)[i]
    }
    return sumaTotal
}

func main(){
    a := []float64{3.00, 3.33, 33.1}

    suma := sumReduce(a)

    fmt.Println(suma)
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

func media(notas map[string]float64)float64{
    sumaTotal := 0.0
    for _, nota := range notas{
        sumaTotal += nota
    }
    return sumaTotal/float64(len(notas))
}

func main(){
    notas := map[string]float64{
        "Carlos" : 7.75,
        "Milagros" : 10.0,
        "Erik" : 5.0,
    }

    fmt.Println(media(notas))
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

func main(){

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

func main(){

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

func main(){

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

func main(){

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

func main(){

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

func main(){

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

func main(){

}
```
