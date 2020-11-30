package e

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
	case B:
		currToken = NextToken(r)
		switch currToken {
		case C:
			currToken = NextToken(r)
			T()
		default:
			panic("Syntax error")
		}
	default:
		panic("Syntax error")
	}
}

func T() {
	switch currToken {
	case A:
		currToken = NextToken(r)
	case F:
		currToken = NextToken(r)
		switch currToken {
		case A:
			currToken = NextToken(r)
		default:
			panic("Syntax error")
		}
	default:
		panic("Syntax error")
	}
}
