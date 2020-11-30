package a

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
	case A:
		currToken = NextToken(r)
		T()
	default:
		panic("Syntax error")
	}
}

func T() {
	switch currToken {
	case A: // First of <S>
		// token is not consumed
		S()
		U()
	}
}

func U() {
	switch currToken {
	case Str, Pl:
		currToken = NextToken(r)
		T()
	default:
		panic("Syntax error")
	}
}
