package i

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
	S()
}

func S() {
	switch currToken {
	case Lpar:
		currToken = NextToken(r)
		T()
	case Rpar, Rbra:
		T()
	default:
		panic("Syntax error")
	}
}

func T() {
	switch currToken {
	case Rpar:
		currToken = NextToken(r)
	case Rbra:
		currToken = NextToken(r)
	default:
		panic("Syntax error")
	}
}
