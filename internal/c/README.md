# Lenguage C

## Gramática

``` text
<L> ::= num | id | (<S>)
<S> ::= <L>{<L>}
```

Que es equivalente a lo siguiente

```text
<S> ::= <L><T>
<T> ::= <L><T> | eps
<L> ::= num | id | (S)
```

solo que la segunda forma tiene la recursión de forma explícita, lo cuál es más
sencillo de implementar.

`num` es una secuencia de dígitos y `id` uns secuencia de letras.
