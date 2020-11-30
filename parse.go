package main

import (
	"bufio"
	"os"
	"psa/d"
)

func main() {
	if len(os.Args) < 2 {
		panic("Input file required")
	}

	file, err := os.Open(os.Args[1])
	if err == nil {
		d.Parse(bufio.NewReader(file))
	}
}
