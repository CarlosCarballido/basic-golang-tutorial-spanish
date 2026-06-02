# Ejercicios por capitulos

Este banco esta organizado segun los 18 temas del tutorial. Cada capitulo incluye una pequena bateria de ejercicios para practicar justo despues de leerlo.

## Capitulo 1. Introduccion

### Ejercicio 1
Enunciado: Crea un programa `main.go` que imprima un saludo, tu nombre y tu carrera.
Pistas: usa `package main`, `import "fmt"` y `fmt.Println`.
Practica: estructura minima de un programa Go.

```go
package main

import "fmt"

func main(){
    nombre := "Carlos"
    edad := 23
    carrera := "Inteligencia Artificial"
    fmt.Printf("Hola Mundo, mi nombre es %s, tengo %d años, estudio %s\n", nombre, edad, carrera)
    fmt.Println("Hola Mundo, mi nombre es", nombre, "tengo", edad, "años", "estudio", carrera)
}
```

### Ejercicio 2
Enunciado: Ejecuta el programa anterior con `go run` y luego genera un binario con `go build`.
Pistas: compara la experiencia de ambos comandos.
Practica: herramientas basicas de Go.

### Ejercicio 3
Enunciado: Escribe un programa que imprima la version de Go que tienes instalada, o al menos el mensaje "Go funciona".
Pistas: usa una salida simple por consola.
Practica: primera ejecucion de un programa.

```go
package main

import (
    "fmt"
    "runtime"
)

func main(){
    version := runtime.Version()
    fmt.Println("Versión instalada de Go:", version)
}
```

## Capitulo 2. Sintaxis basica

### Ejercicio 4
Enunciado: Crea un programa que use `if`, `else if` y `else` para clasificar un numero como negativo, cero o positivo.
Pistas: define una variable local y comparala con 0.
Practica: condicionales.

```go
package main

import "fmt"

func main(){
    num := 1
    if num < 0{
        fmt.Println("Es negativo")
    } else if num == 0{
        fmt.Println("Es 0")
    } else {
        fmt.Println("Es positivo")
    }
}
```

### Ejercicio 5
Enunciado: Escribe tres versiones del mismo recorrido: un `for` clasico, un `for` tipo `while` y un `for` infinito con `break`.
Pistas: usa el mismo rango de valores en cada version.
Practica: sintaxis de `for`.

```go
package main

import "fmt"

func main(){
    for i:=0; i < 3; i++ {
        fmt.Println(i)
    }

    i := 0

    for i < 3 {
        fmt.Println(i)
        i++
    }

    j := 0

    for {
        if j >= 3 {
            break
        }
        fmt.Println(j)
        j++
    }
}
```

### Ejercicio 6
Enunciado: Crea un `switch` que reciba un dia de la semana y muestre si es laborable o fin de semana.
Pistas: agrupa varios casos en una sola rama.
Practica: `switch` y lectura de casos.

```go
package main

import "fmt"

func main(){
    dia := "martes"

    switch dia {
        case "lunes":
            fmt.Println("Laboral")
        case "martes":
            fmt.Println("Laboral")
        case "miercoles":
            fmt.Println("Laboral")
        case "jueves":
            fmt.Println("Laboral")
        case "viernes":
            fmt.Println("Laboral")
        case "sabado":
            fmt.Println("Fin de semana")
        case "domingo":
            fmt.Println("Fin de semana")
    }

    switch dia {
        case "lunes", "martes", "miercoles", "jueves", "viernes":
            fmt.Println("Laboral")
        case "sabado", "domingo":
            fmt.Println("Fin de semana")
    }
}
```

## Capitulo 3. Tipos, variables y constantes

### Ejercicio 7
Enunciado: Declara una variable con `var`, otra con `:=` y una constante; imprime sus valores.
Pistas: usa tipos distintos, por ejemplo `string`, `int` y `float64`.
Practica: declaracion de variables y constantes.

```go
package main

import "fmt"

func main() {
    var nombre string = "Carlos"
    const colorOjos string = "verde"
    nombreNoTipado := "Carlos"
    var letra rune = 'C'
    var edad int = 23
    var altura float64 = 1.83

    fmt.Println("Nombre:", nombre)
    fmt.Println("Nombre no tipado:", nombreNoTipado)
    fmt.Println("Letra:", string(letra))
    fmt.Println("Edad:", edad)
    fmt.Println("Altura:", altura)
    fmt.Println("Color de ojos:", colorOjos)
}

```

### Ejercicio 8
Enunciado: Convierte una variable entera a `float64` y calcula su promedio con otro valor decimal.
Pistas: la conversion debe ser explicita.
Practica: conversion de tipos.

```go
package main

import "fmt"

func main() {
    var edad int = 23
    var altura float64 = 1.83

    fmt.Println("Edad:", edad)
    fmt.Println("Altura:", altura)

    edadFloat := float64(edad)

    promedio := (edadFloat + altura) / 2

    fmt.Println("Edad convertida:", edadFloat)
    fmt.Println("Promedio:", promedio)
}
```

### Ejercicio 9
Enunciado: Crea un programa que cuente cuantas vocales tiene una palabra.
Pistas: recorre la cadena y compara cada caracter.
Practica: `string`, `rune` y control de flujo.

```go
package main

import "fmt"

func main() {
    var palabra string = "palabra"
    var caracter rune
    var contador int = 0

    for i:=0; i < len(palabra); i++{
        caracter = rune(palabra[i])
        fmt.Println(string(caracter))
        if caracter == 'a' ||
            caracter == 'e' ||
            caracter == 'i' ||
            caracter == 'o' ||
            caracter == 'u' {
            contador++
        }
    }
    fmt.Println("Número de vocales:", contador)
}
```

## Capitulo 4. Arrays, slices y maps

### Ejercicio 10
Enunciado: Define un array de 4 enteros, rellena sus posiciones e imprime su contenido.
Pistas: el tamano debe ser fijo.
Practica: arrays.

```go
package main

import "fmt"

func main() {
    
    // alternativa array := [4]int{1, 2, 3, 4}
    var array [4]int
    array[0] = 1
    array[1] = 2
    array[2] = 3
    array[3] = 4

    // alternativa: for i, valor := range array
    for i:=0; i<len(array); i++ {
        fmt.Println("posicion:", i, "del array:", array[i])
    }
}
```

### Ejercicio 11
Enunciado: Crea un slice de materias, añade una materia nueva con `append` y muestra la longitud final.
Pistas: reasigna el resultado de `append`.
Practica: slices dinamicos.

```go
package main

import "fmt"

func main() {
    materias := []string{"MITEX", "APAUBIO", "CCPD", "IASA"}
    materia := "SIRE"

    fmt.Println("Materias:", materias)
    fmt.Println("Longitud:", len(materias))

    fmt.Println("Realizamos append")
    materias = append(materias, materia)

    fmt.Println("Materias:", materias)
    fmt.Println("Longitud:", len(materias))
}
```

### Ejercicio 12
Enunciado: Usa un `map[string]int` para guardar el numero de creditos de varias asignaturas y consulta una de ellas.
Pistas: añade una comprobacion de existencia.
Practica: maps.

```go
package main

import "fmt"

func main() {
    materias := map[string]int{
        "MITEX" : 6,
         "APAUBIO" : 9,
         "CCPD" : 9,
         "IASA" : 3,
         "SIRE" : 6
    }

    // comprobar existencia
    creditos, existe := materias["SIRE"]

    if existe {
        fmt.Println("SIRE tiene", creditos, "créditos")
    } else {
        fmt.Println("La asignatura no existe")
    }
}
```


## Capitulo 5. Structs y composicion

### Ejercicio 13
Enunciado: Define un `struct` `Libro` con titulo, autor y paginas, y crea una instancia de ejemplo.
Pistas: usa un literal de struct.
Practica: modelado de datos.

```go
package main

import "fmt"

type Libro struct{
    Titulo string
    Autor string
    Paginas int
}

func (l Libro) getAutor() string{
    return l.Autor
}

func (l Libro) getTitulo() string{
    return l.Titulo
}

func (l *Libro) SetPaginas(nPag int) {
    l.Paginas = nPag
}

func main() {
    libro := Libro{
        Titulo:  "El Quijote",
        Autor:   "Miguel de Cervantes",
        Paginas: 863,
    }

    fmt.Println("Título:", libro.GetTitulo())
    fmt.Println("Autor:", libro.GetAutor())
    fmt.Println("Páginas:", libro.Paginas)

    libro.SetPaginas(900)

    fmt.Println("Páginas:", libro.Paginas)
}
```


### Ejercicio 14
Enunciado: Crea un `struct` `Alumno` que embeba otro `struct` `Persona` y accede a un campo promocionado.
Pistas: usa composicion por embebido.
Practica: composicion.

```go
package main

import "fmt"

type Persona struct{
    Nombre string
}


type Alumno struct{
    Persona
    Curso string
}

func main() {
    Carlos := Alumno{
        Persona{
            Nombre: "Carlos",
        },
        Curso: "4 GRIA",
    }

    fmt.Println("Nombre:", Carlos.Nombre)
    fmt.Println("Autor:", Carlos.Curso)
}
```


### Ejercicio 15
Enunciado: Implementa una funcion `ResumenPedido` que reciba un `struct` con varios campos y devuelva un texto legible.
Pistas: usa `fmt.Sprintf`.
Practica: structs y funciones auxiliares.

```go
package main

import "fmt"

type Pedido struct {
	Cliente string
	Producto string
	Cantidad int
	Precio float64
}

func ResumenPedido(p Pedido) string {
	total := float64(p.Cantidad) * p.Precio

	return fmt.Sprintf(
		"Cliente: %s | Producto: %s | Cantidad: %d | Precio unitario: %.2f€ | Total: %.2f€",
		p.Cliente,
		p.Producto,
		p.Cantidad,
		p.Precio,
		total,
	)
}

func main() {
	pedido := Pedido{
		Cliente: "Carlos",
		Producto: "Monitor",
		Cantidad: 2,
		Precio: 199.99,
	}

	resumen := ResumenPedido(pedido)

	fmt.Println(resumen)
}
```

## Capitulo 6. Punteros

### Ejercicio 16
Enunciado: Escribe una funcion que reciba un puntero a entero y multiplique su valor por dos.
Pistas: recuerda desreferenciar con `*`.
Practica: punteros y mutacion.

```go
package main

import "fmt"

func multi(num *int){
    *num = *num * 2
}

func main() {
    valor := 5
    multi(&valor)
}
```


### Ejercicio 17
Enunciado: Crea una funcion que reciba un `struct` por puntero y actualice uno de sus campos.
Pistas: compara el valor antes y despues de llamar a la funcion.
Practica: paso por referencia.

```go
package main

import "fmt"

type Persona struct {
	Nombre string
	Altura int
}

func actualizarAltura(p *Persona, nuevaAltura int) {
	p.Altura = nuevaAltura
}

func main() {
	p := Persona{
		Nombre: "Carlos",
		Altura: 181,
	}

	fmt.Println("Antes:", p.Altura)

	actualizarAltura(&p, 185)

	fmt.Println("Después:", p.Altura)
}
```


### Ejercicio 18
Enunciado: Protege una posible desreferencia de puntero `nil` con una comprobacion previa.
Pistas: usa `if p != nil`.
Practica: seguridad al trabajar con punteros.

```go
package main

import "fmt"

func main() {
    var p *int
    if p != nil {
        fmt.Println(*p)
    }
}
```


## Capitulo 7. Funciones y metodos

### Ejercicio 19
Enunciado: Escribe una funcion que devuelva el resultado de una division y un error si el divisor es cero.
Pistas: usa retorno multiple.
Practica: funciones y errores.

```go
package main

import (
    "fmt"
    )

func dividir(a int, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division por cero")
	}

	return float64(a) / float64(b), nil
}

func main() {
	resultado, err := dividir(100, 10)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(resultado)
}
```


### Ejercicio 20
Enunciado: Crea un `struct` `Cuenta` y define un metodo con receptor por puntero para ingresar dinero.
Pistas: el metodo debe modificar el saldo.
Practica: metodos con receptor por puntero.

```go
package main

import "fmt"

type Cuenta struct {
	Saldo float64
}

func (c *Cuenta) Ingreso(cantidad float64) {
	c.Saldo += cantidad
}

func main() {
	c1 := Cuenta{Saldo: 10.0}

	fmt.Println("Antes:", c1.Saldo)

	c1.Ingreso(20.0)

	fmt.Println("Después:", c1.Saldo)
}
```


### Ejercicio 21
Enunciado: Define dos metodos sobre un mismo tipo: uno que solo lea datos y otro que los modifique.
Pistas: compara receptor por valor y por puntero.
Practica: diseno de metodos.

```go
package main

import "fmt"

type Cuenta struct {
	Saldo float64
}

func (c Cuenta) GetSaldo() float64 {
	return c.Saldo
}

func (c *Cuenta) SetSaldo(saldo float64) {
	c.Saldo = saldo
}

func main() {
	saldo := 10.0

	c1 := Cuenta{Saldo: saldo}

	fmt.Println("Saldo inicial:", c1.GetSaldo())

	c1.SetSaldo(25.0)

	fmt.Println("Saldo final:", c1.GetSaldo())
}
```


## Capitulo 8. Paquetes e interfaces

### Ejercicio 22
Enunciado: Crea un paquete propio con una funcion exportada y otra privada.
Pistas: usa mayuscula para exportar.
Practica: organizacion en paquetes.

```go
package utilidades

import "fmt"

// Función exportada (pública)
func Saludar(nombre string) {
	fmt.Println(obtenerMensaje(nombre))
}

// Función privada
func obtenerMensaje(nombre string) string {
	return "Hola " + nombre
}
```


```go
package main

import "miProyecto/utilidades"

func main() {
	utilidades.Saludar("Carlos")
}
```


### Ejercicio 23
Enunciado: Define una interfaz `Notificador` y dos tipos que la implementen de forma implicita.
Pistas: ambos tipos deben tener el mismo metodo.
Practica: interfaces.

```go
package notificaciones

import "fmt"

type Notificador interface() {
    
}
```

```go
package main

import "miProyecto/notifcaciones"

func main() {
	notificaciones.Notificar("Carlos")
}
```


### Ejercicio 24
Enunciado: Crea una funcion que reciba una interfaz y la use sin conocer el tipo concreto.
Pistas: pasa un valor que cumpla el contrato.
Practica: polimorfismo por interfaces.

```go
package main

import "fmt"

// Interfaz
type Notificador interface {
	Notificar()
}

// Tipo 1
type Email struct {
	Destinatario string
}

func (e Email) Notificar() {
	fmt.Println("Enviando email a:", e.Destinatario)
}

// Tipo 2
type SMS struct {
	Numero string
}

func (s SMS) Notificar() {
	fmt.Println("Enviando SMS a:", s.Numero)
}

// Función que recibe la interfaz
func Enviar(n Notificador) {
	n.Notificar()
}

func main() {
	email := Email{
		Destinatario: "carlos@email.com",
	}

	sms := SMS{
		Numero: "600123123",
	}

	Enviar(email)
	Enviar(sms)
}
```


## Capitulo 9. Fundamentos de concurrencia

### Ejercicio 25
Enunciado: Explica por escrito la diferencia entre concurrencia, paralelismo y asincronia usando un ejemplo cotidiano.
Pistas: piensa en cocina, transporte o atencion al cliente.
Practica: razonamiento conceptual.

Diferencia entre concurrencia, paralelismo y asincronía

Imagina una cocina:

Concurrencia

Hay un solo cocinero preparando varios platos.

Preparar ensalada
↓
Esperar ingredientes
↓
Preparar sopa
↓
Remover ensalada
↓
Servir sopa

El cocinero alterna tareas. No hace varias cosas al mismo tiempo, pero todas avanzan.

Paralelismo

Hay dos cocineros.

Cocinero 1 -> ensalada
Cocinero 2 -> sopa

Ambos trabajan exactamente al mismo tiempo.

Asincronía

El cocinero mete una pizza en el horno y sigue trabajando.

Meter pizza al horno
↓
La pizza se cocina sola
↓
Mientras tanto preparo una ensalada
↓
La pizza termina

La tarea continúa sin bloquear al cocinero.

Resumen:

Concurrencia -> varias tareas progresan intercaladas.
Paralelismo  -> varias tareas ejecutándose simultáneamente.
Asincronía   -> una tarea espera mientras otra continúa.

### Ejercicio 26
Enunciado: Dibuja una linea temporal ASCII con dos tareas intercaladas y otra en paralelo.
Pistas: usa bloques y flechas.
Practica: representacion visual de concurrencia.

Línea temporal ASCII
Concurrencia

Dos tareas avanzan intercaladas:

Tiempo ──────────────────────────>

Tarea A: [Trabaja]-----[Trabaja]-----
Tarea B: -----[Trabaja]-----[Trabaja]
Paralelismo

Dos tareas ejecutándose simultáneamente:

Tiempo ──────────────────────────>

Tarea A: [Trabaja][Trabaja][Trabaja]
Tarea B: [Trabaja][Trabaja][Trabaja]
Asincronía

Una tarea espera mientras otra continúa:

Tiempo ──────────────────────────>

Tarea A: [Lanza tarea]----[Sigue trabajando]----
Tarea B: ----[Procesando]----[Termina]----------

### Ejercicio 27
Enunciado: Identifica si tres escenarios concretos son concurrentes, paralelos o asincronos.
Pistas: redacta una frase por escenario.
Practica: analisis conceptual.

Escenario 1

Un camarero atiende tres mesas distintas alternando entre ellas.

Respuesta: Concurrencia.

Porque una sola persona gestiona varias tareas que avanzan de forma intercalada.

Escenario 2

Dos empleados atienden clientes diferentes al mismo tiempo en dos ventanillas.

Respuesta: Paralelismo.

Porque las tareas se ejecutan simultáneamente.

Escenario 3

Envías un correo electrónico y continúas escribiendo un informe mientras se envía.

Respuesta: Asincronía.

Porque la operación de envío continúa sin bloquear tu trabajo.

## Capitulo 10. Goroutines

### Ejercicio 28
Enunciado: Lanza una goroutine que imprima un mensaje y espera a que termine.
Pistas: usa `WaitGroup` o una pequena sincronizacion equivalente.
Practica: creacion y ciclo de vida de goroutines.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func mesajeAndSleep(wg *sync.WaitGroup){
    defer wg.Done()

    fmt.Println("Hola")
    time.Sleep(100 * time.Millisecond)
    fmt.Println("Adios")
}

func main(){
    var wg sync.WaitGroup
    for i:=1; i<5; i++ {
        wg.Add(1)
        go mesajeAndSleep(&wg)
    }
    wg.Wait()
}

```

### Ejercicio 29
Enunciado: Lanza cinco goroutines dentro de un bucle y pasa el indice correctamente a cada una.
Pistas: evita capturar la variable del bucle sin copiarla.
Practica: goroutines y closures.

```go
package main

import (
    "fmt"
    "sync"
)

func Saludo(nombre string, wg *sync.WaitGroup){
    defer wg.Done()

    fmt.Println("Hola", nombre)
}

func main(){
    var wg sync.WaitGroup
    for i:=0; i <= 5; i++{
        wg.Add(1)
        go Saludo("Carlos", &wg)
    }
    wg.Wait()
}

```

### Ejercicio 30
Enunciado: Simula cien tareas concurrentes y cuenta cuantas terminan.
Pistas: coordina el final con `WaitGroup`.
Practica: escalado de goroutines.

```go
package main

import (
	"fmt"
	"sync"
)

func doWork(wg *sync.WaitGroup, contador *int, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	*contador++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	contador := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doWork(&wg, &contador, &mu)
	}

	wg.Wait()

	fmt.Println("Tareas completadas:", contador)
}

```

## Capitulo 11. Channels

### Ejercicio 31
Enunciado: Crea un channel sin buffer y haz pasar un valor de una goroutine a otra.
Pistas: separa productor y consumidor.
Practica: envio y recepcion.

### Ejercicio 32
Enunciado: Construye un canal de resultados, cierralo correctamente y recorre sus valores con `range`.
Pistas: el cierre debe hacerlo quien produce.
Practica: cierre de channels.

### Ejercicio 33
Enunciado: Implementa un ejemplo productor-consumidor con dos goroutines y un channel.
Pistas: piensa en una tuberia simple.
Practica: coordinacion basica con channels.

## Capitulo 12. Select

### Ejercicio 34
Enunciado: Usa `select` para recibir de dos channels distintos y mostrar el primero que llegue.
Pistas: no uses `default` al principio.
Practica: multiplexacion.

### Ejercicio 35
Enunciado: Agrega un timeout con `time.After` a un `select`.
Pistas: define cuanto tiempo esperas antes de cancelar.
Practica: timeouts.

### Ejercicio 36
Enunciado: Usa `default` en un `select` para evitar bloqueo cuando no hay mensajes.
Pistas: muestra un mensaje alternativo.
Practica: comportamiento no bloqueante.

## Capitulo 13. Buffered channels

### Ejercicio 37
Enunciado: Crea un channel con capacidad 2 y comprueba que los dos primeros envios no bloquean.
Pistas: lee despues de enviar.
Practica: capacidad y bloqueo.

### Ejercicio 38
Enunciado: Cambia la capacidad del channel anterior y observa cuando aparece el bloqueo.
Pistas: prueba con 0, 1 y 3.
Practica: efecto del buffer.

### Ejercicio 39
Enunciado: Escribe un caso en el que un buffered channel sea util y otro en el que no lo sea.
Pistas: razonalo a nivel de diseno.
Practica: criterio de uso.

## Capitulo 14. Patrones de concurrencia

### Ejercicio 40
Enunciado: Implementa un worker pool con tres workers y cinco tareas.
Pistas: separa `jobs` y `results`.
Practica: worker pool.

### Ejercicio 41
Enunciado: Implementa fan-in uniendo resultados de dos workers en un canal comun.
Pistas: crea una funcion `merge`.
Practica: fan-in.

### Ejercicio 42
Enunciado: Construye un pipeline de dos etapas que transforme texto y lo envie a una salida final.
Pistas: cada etapa debe hacer una sola cosa.
Practica: pipeline.

## Capitulo 15. Sincronizacion

### Ejercicio 43
Enunciado: Protege un contador compartido con `Mutex`.
Pistas: toda escritura debe quedar dentro de la seccion critica.
Practica: exclusión mutua.

### Ejercicio 44
Enunciado: Crea una cache de lectura frecuente con `RWMutex`.
Pistas: usa `RLock` para lectura y `Lock` para escritura.
Practica: RWMutex.

### Ejercicio 45
Enunciado: Lanza varias goroutines y espera a que todas terminen con `WaitGroup`.
Pistas: usa `Add`, `Done` y `Wait`.
Practica: coordinacion de finalizacion.

## Capitulo 16. Context

### Ejercicio 46
Enunciado: Cancela una goroutine con `context.WithCancel`.
Pistas: escucha `ctx.Done()` en el worker.
Practica: cancelacion cooperativa.

### Ejercicio 47
Enunciado: Aborta una tarea que tarda demasiado usando `context.WithTimeout`.
Pistas: devuelve un mensaje distinto si vence el tiempo.
Practica: timeout.

### Ejercicio 48
Enunciado: Propaga un mismo context por tres funciones encadenadas.
Pistas: pasa `context.Context` como primer parametro.
Practica: propagacion de contexto.

## Capitulo 17. Errores comunes

### Ejercicio 49
Enunciado: Revisa un fragmento con una posible race condition y explica por que es inseguro.
Pistas: busca variables compartidas sin proteccion.
Practica: deteccion de carreras.

### Ejercicio 50
Enunciado: Localiza un deadlock en un ejemplo de channels y describe como evitarlo.
Pistas: analiza quien envia y quien recibe.
Practica: diagnostico de bloqueos.

### Ejercicio 51
Enunciado: Corrige una goroutine leak provocada por un canal que nunca se cierra.
Pistas: decide quien produce y quien cierra.
Practica: cierre correcto y limpieza.

## Capitulo 18. Proyecto practico

### Ejercicio 52
Enunciado: Diseña el esquema de modulos para un procesador concurrente de tareas.
Pistas: separa productor, workers, agregador y estadisticas.
Practica: arquitectura del proyecto.

### Ejercicio 53
Enunciado: Añade cancelacion por timeout global al procesador de tareas.
Pistas: combina `context` y `select`.
Practica: control temporal del sistema.

### Ejercicio 54
Enunciado: Extiende el proyecto para contar tareas completadas, fallidas y canceladas.
Pistas: protege las estadisticas con `Mutex`.
Practica: integracion de sincronizacion y observabilidad.
