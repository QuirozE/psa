package j

import (
	"bufio"
	"fmt"
)

var r *bufio.Reader
var currToken Token

func willAccept() {
	if r := recover(); r != nil {
		fmt.Println("Reject")
	} else if currToken != Eof {
		fmt.Println("Reject")
	} else {
		fmt.Println("Accept")
	}
}

func Parse(reader *bufio.Reader) {
	defer willAccept()
	r = reader
	currToken = NextToken(r)
	E()
}

func E() {
	F()
	J()
	K()
}

func K() {
	switch currToken {
	case Mod:
		currToken = NextToken(r)
		F()
		J()
		K()
	}
}

func J() {
	switch currToken {
	case Not, Lpar, A, B:
		F()
		J()
	}
}

func F() {
	switch currToken {
	case Not:
		F()
		H()
	case Lpar, A, B:
		G()
		H()
	default:
		panic("Syntax error")
	}
}

func H() {
	switch currToken {
	case Star, Option:
		currToken = NextToken(r)
	}
}

func G() {
	switch currToken {
	case Lpar:
		currToken = NextToken(r)
		E()
		switch currToken {
		case Rpar:
			currToken = NextToken(r)
		default:
			panic("Syntax error")
		}
	case A, B:
		currToken = NextToken(r)
	default:
		panic("Syntax error")
	}
}
