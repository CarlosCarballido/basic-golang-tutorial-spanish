# Ejercicios avanzados

## Ejercicio 33
Enunciado: Lanza una goroutine que imprima un mensaje y espera correctamente a que termine.
Pistas: usa `WaitGroup`.
Practica: goroutines, sincronizacion.

## Ejercicio 34
Enunciado: Crea dos goroutines que envien datos a un channel y recoge los resultados.
Pistas: piensa en productor-consumidor.
Practica: channels, concurrencia.

## Ejercicio 35
Enunciado: Implementa un `select` con un timeout de un segundo.
Pistas: usa `time.After`.
Practica: select, timeout.

## Ejercicio 36
Enunciado: Crea un buffered channel y explica por que no bloquea en los primeros envios.
Pistas: prueba diferentes capacidades.
Practica: buffered channels.

## Ejercicio 37
Enunciado: Implementa un worker pool con tres workers y cinco tareas.
Pistas: separa jobs y results.
Practica: patron worker pool.

## Ejercicio 38
Enunciado: Protege un contador compartido con `Mutex`.
Pistas: revisa cada acceso al dato.
Practica: mutex, race conditions.

## Ejercicio 39
Enunciado: Diseña una cache concurrente de solo lectura frecuente con `RWMutex`.
Pistas: distingue lector y escritor.
Practica: RWMutex.

## Ejercicio 40
Enunciado: Usa `context.WithCancel` para detener un worker.
Pistas: escucha `ctx.Done()`.
Practica: context, cancelacion.

## Ejercicio 41
Enunciado: Usa `context.WithTimeout` para abortar una tarea lenta.
Pistas: responde al timeout.
Practica: context, timeout.

## Ejercicio 42
Enunciado: Construye un pipeline de dos etapas con channels.
Pistas: una etapa transforma y otra agrega.
Practica: pipeline.

## Ejercicio 43
Enunciado: Implementa fan-out repartiendo tareas entre varias goroutines.
Pistas: comparte la entrada.
Practica: fan-out.

## Ejercicio 44
Enunciado: Implementa fan-in unificando salidas de varias goroutines.
Pistas: usa un canal de merge.
Practica: fan-in.

## Ejercicio 45
Enunciado: Detecta y corrige una race condition en un contador global.
Pistas: ejecuta mentalmente varias intercalaciones.
Practica: diagnostico de concurrencia.

## Ejercicio 46
Enunciado: Evita una goroutine leak en un consumidor de channels.
Pistas: define cierre y cancelacion.
Practica: leaks, cierre correcto.

## Ejercicio 47
Enunciado: Reescribe una solucion basada en estado compartido usando channels.
Pistas: compara claridad de ambas versiones.
Practica: diseno concurrente.

## Ejercicio 48
Enunciado: Crea un sistema que procese tareas y pueda ser cancelado desde fuera.
Pistas: combina context y select.
Practica: cancelacion cooperativa.

## Ejercicio 49
Enunciado: Simula una operacion remota con timeout y manejo de error.
Pistas: no bloquees indefinidamente.
Practica: context, select, errores.

## Ejercicio 50
Enunciado: Añade estadisticas protegidas por mutex a un worker pool.
Pistas: separa datos compartidos y flujo de mensajes.
Practica: mutex, patrones.

## Ejercicio 51
Enunciado: Implementa un lector de eventos con `select` que cierre correctamente al terminar.
Pistas: evita canales colgantes.
Practica: select, shutdown limpio.

## Ejercicio 52
Enunciado: Diseña una solucion concurrente para calcular cuadrados de muchos numeros sin saturar recursos.
Pistas: worker pool o fan-out controlado.
Practica: escalado, limitacion.
