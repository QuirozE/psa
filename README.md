# Análsis sintáctico descendente recursivo

Como parte de una tarea, hay que implementar un analizador recursivo para diez
gramáticas diferentes.

Aquí se incluye el código fuente de esos analizadores.

## Uso

```console
.$ go build
./psa
Usage: ./psa <grammar> <file>

where
<grammar> ::= a | b | c | d | e | f | g | h | i | j | k

and <file> is the source file.
```

En la carpeta `examples` hay algunos archivos válidos para cada gramática.

## Gramáticas

La especificación de cada gramática se puede encontrar en el README de su
paquete correspondiente.
