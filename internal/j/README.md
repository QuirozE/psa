# Lenguage J

## Gramática

```text
<E> ::= <F><J><K>
<K> ::= {%<F><J>}
<J> ::= {F}
<F> ::= ~<F><H> | <G><H>
<H> ::= * | ? | eps
<G> ::= (<E>) | a | b
```
