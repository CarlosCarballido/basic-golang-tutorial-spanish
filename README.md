# Tutorial completo de Go

Este repositorio contiene un tutorial progresivo y autocontenido de Go pensado para estudiantes universitarios que empiezan desde cero y quieren llegar a concurrencia intermedia-avanzada.

## Como usarlo

Lee los capitulos en orden. Los capitulos 1 a 8 cubren la base del lenguaje. Los capitulos 9 a 17 se centran en concurrencia, coordinacion y patrones reales. El capitulo 18 integra todo en un proyecto practico.

## Estructura

- [01-introduccion.md](01-introduccion.md)
- [02-sintaxis-basica.md](02-sintaxis-basica.md)
- [03-tipos-y-variables.md](03-tipos-y-variables.md)
- [04-arrays-slices-maps.md](04-arrays-slices-maps.md)
- [05-structs-y-composicion.md](05-structs-y-composicion.md)
- [06-punteros.md](06-punteros.md)
- [07-funciones-y-metodos.md](07-funciones-y-metodos.md)
- [08-paquetes-e-interfaces.md](08-paquetes-e-interfaces.md)
- [09-concurrencia-fundamentos.md](09-concurrencia-fundamentos.md)
- [10-goroutines.md](10-goroutines.md)
- [11-channels.md](11-channels.md)
- [12-select.md](12-select.md)
- [13-buffered-channels.md](13-buffered-channels.md)
- [14-patrones-de-concurrencia.md](14-patrones-de-concurrencia.md)
- [15-sincronizacion.md](15-sincronizacion.md)
- [16-context.md](16-context.md)
- [17-errores-comunes.md](17-errores-comunes.md)
- [18-proyecto-practico.md](18-proyecto-practico.md)

## Carpetas de codigo

- [codigo/ejemplos-basicos](codigo/ejemplos-basicos)
- [codigo/concurrencia](codigo/concurrencia)
- [codigo/channels](codigo/channels)
- [codigo/sincronizacion](codigo/sincronizacion)
- [codigo/proyecto-final](codigo/proyecto-final)

## Ejercicios

- [ejercicios/ejercicios-basicos.md](ejercicios/ejercicios-basicos.md)
- [ejercicios/ejercicios-intermedios.md](ejercicios/ejercicios-intermedios.md)
- [ejercicios/ejercicios-avanzados.md](ejercicios/ejercicios-avanzados.md)
- [ejercicios/ejercicios-por-capitulos.md](ejercicios/ejercicios-por-capitulos.md)
- [ejercicios/soluciones.md](ejercicios/soluciones.md)

El archivo por capitulos agrupa un conjunto de ejercicios para cada tema explicado en el tutorial.

## Idea pedagogica

Cada capitulo combina teoria, analogias, diagramas ASCII, ejemplos completos, explicacion paso a paso, errores frecuentes, buenas practicas, comparacion con Java y preguntas de autoevaluacion.

## Proyecto final

El capitulo final desarrolla un procesador concurrente de tareas con `goroutines`, `channels`, `select`, `WaitGroup`, `Mutex` y `context`.
