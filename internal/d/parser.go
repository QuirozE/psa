package d

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
	case Id, Num, Lpar:
		L()
		T()
	default:
		panic("Syntax error")
	}
}

func T() {
	switch currToken {
	case Comma:
		currToken = NextToken(r)
		S()
	}
}

func L() {
	switch currToken {
	case Num, Id:
		currToken = NextToken(r)
	case Lpar:
		currToken = NextToken(r)
		S()
		switch currToken {
		case Rpar:
			currToken = NextToken(r)
		default:
			panic("Syntax error")
		}
	default:
		panic("Syntax error")
	}
}
