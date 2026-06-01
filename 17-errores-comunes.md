# 17. Errores comunes en Go concurrente

Este capitulo resume los fallos que mas se repiten cuando un programa concurrente parece correcto pero falla en produccion o en pruebas.

## 1. Race conditions

Varios goroutines acceden y modifican el mismo dato sin proteccion.

## 2. Deadlocks

Dos o mas goroutines se esperan mutuamente y nadie avanza.

## 3. Goroutine leaks

Una goroutine queda viva para siempre porque nadie la cancela o consume su canal.

## 4. Cierre incorrecto de channels

Solo el productor deberia cerrar. Cerrar dos veces o cerrar desde el consumidor suele causar panics o diseno confuso.

## 5. Captura incorrecta de variables de bucle

```go
for i := 0; i < 3; i++ {
    go func() {
        fmt.Println(i)
    }()
}
```

La solucion es pasar `i` como argumento al closure.

## 6. Nil channels

Un channel `nil` bloquea para siempre si se usa en envio o recepcion.

## 7. Copiar mutex o WaitGroup

Los tipos de sincronizacion no deben copiarse una vez en uso.

## 8. Ignorar errores

En concurrencia, ignorar un error puede esconder una cancelacion, un timeout o un fallo de integracion.

## 9. Compartir estado cuando un channel bastaba

Si el problema se expresa mejor como flujo de mensajes, un channel suele ser mas claro que un estado global protegido.

## 10. No cerrar recursos

Canales, goroutines, archivos y contexts deben cerrarse o cancelarse cuando termina la operacion.

## Lista de diagnostico

- Hay una sola fuente de verdad?
- Quien produce y quien consume?
- Quien cancela?
- Quien cierra?
- Hay proteccion suficiente para el estado compartido?

## Comparacion con Java

- En Java muchos errores analogos aparecen con `ExecutorService`, locks o colas concurrentes.
- Go no elimina los problemas; solo los hace mas explicitos y, por tanto, mas faciles de razonar si se usa bien.

## Preguntas de autoevaluacion

1. Que sintomas tiene un deadlock?
2. Como identificar una goroutine leak?
3. Por que la captura de variables de bucle es tan peligrosa?

## Ejercicio guiado

Lee un fragmento con goroutines y explica donde podria aparecer una race condition.

## Ejercicio sin resolver

Toma un programa concurrente simple y enumera tres fallos posibles si se ejecuta con carga alta.
