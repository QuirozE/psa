# Lenguage C

## Gramática
L→num | id | (S)
S→L{L}

```text
<S> ::= <L><T>
<T> ::= ,<S> | eps
<L> ::= num | id | (S)
```

where `num` is a sequence of digits and `id` is a sequence of letters.
