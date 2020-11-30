package b

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
	case Zero:
		currToken = NextToken(r)
		T()
	default:
		panic("Syntax error")
	}
}

func T() {
	switch currToken {
	case Zero: // First of [<S>]
		// token is not consumed
		S()
		switch currToken {
		case One:
			currToken = NextToken(r)
		default:
			panic("Syntax error")
		}
	case One:
		// consume token and end execution
		currToken = NextToken(r)
	default:
		panic("Syntax error")
	}
}
