# Práctica 2: Cifrado de Vigenere

**Objetivo**: Implementar el cifrado de Vigenere.

**Desarrollo**:

Implementa el cifrado y descifrado de Vigenere según la descripción que se incluye a continuación:

Se usará el alfabeto sin Ñ, con W y sin espacios, luego el módulo para las operaciones será m=26.

El cifrado opera sobre bloques de letras, y la clave es una palabra o frase que se repite cuantas veces sea necesario.

Si la clave tiene longitud *r*, entonces el texto se divide en bloques de longitud r y la clave se suma a cada bloque, sumando letra a letra módulo 26, para producir el texto cifrado.

Dicho de otra forma, dada la clave k<sub>1</sub> k<sub>2</sub>...k<sub>r</sub> introducida por el usuario, la primera letra del texto original introducido por el usuario se sustituye por otra que ocupa k<sub>1</sub> posiciones más allá en el alfabeto, la segunda por la que ocupa k<sub>2</sub> posiciones más allá, y así sucesivamente.

**Ejemplo:**

* Palabra clave: `MISION`
* Texto original: `ESTE MENSAJE SE AUTODESTRUIRA`
* Mezclado:

  |          |          |          |          |      |
  |----------|----------|----------|----------|------|
  | `ESTEME` | `NSAJES` | `EAUTOD` | `ESTRUI` | `RA` |
  | `MISION` | `MISION` | `MISION` | `MISION` | `MI` |
  | `QALMAR` | `ZASRSF` | `QIMBCQ` | `QALIZV` | `DI` |

* Texto cirfrado: `QALMARZASRSFQIMBCQQALIZVDI`

## Desarrollo

Para ejecutar las pruebas:

```
ginkgo './...'
```

Para ejecutar el programa principal

```
go run ./cmd/main.go
```
