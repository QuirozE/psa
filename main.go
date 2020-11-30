package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/QuirozE/psa/internal/a"
	"github.com/QuirozE/psa/internal/b"
	"github.com/QuirozE/psa/internal/c"
	"github.com/QuirozE/psa/internal/d"
	"github.com/QuirozE/psa/internal/e"
	"github.com/QuirozE/psa/internal/f"
	"github.com/QuirozE/psa/internal/g"
	"github.com/QuirozE/psa/internal/h"
	"github.com/QuirozE/psa/internal/i"
	"github.com/QuirozE/psa/internal/j"
)

var use = `
Usage: ./psa <grammar> <file>

where
<grammar> ::= a | b | c | d | e | f | g | h | i | j | k

and <file> is the source file.
`

func usage() {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println(use)
	}
}

func main() {
	defer usage()

	if len(os.Args) < 2 {
		panic("Missing grammar selection")
	}

	pSel := os.Args[1]

	if len(os.Args) < 3 {
		panic("Input file required")
	}

	fileName := os.Args[2]

	file, err := os.Open(fileName)

	parsers := map[string]func(*bufio.Reader){
		"a": a.Parse,
		"b": b.Parse,
		"c": c.Parse,
		"d": d.Parse,
		"e": e.Parse,
		"f": f.Parse,
		"g": g.Parse,
		"h": h.Parse,
		"i": i.Parse,
		"j": j.Parse,
	}

	if err == nil {
		if _, ok := parsers[pSel]; ok {
			parsers[pSel](bufio.NewReader(file))
		} else {
			panic("Invalid grammar")
		}
	}
}
