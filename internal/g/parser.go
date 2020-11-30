package g

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
	case Id:
		currToken = NextToken(r)
		switch currToken {
		case Ddot:
			currToken = NextToken(r)
			R()
			T()
		default:
			panic("Syntax error")
		}
	default:
		panic("Syntax error")
	}
}

func R() {
	switch currToken {
	case Id:
		currToken = NextToken(r)
		R()
	}
}

func T() {
	switch currToken {
	case Id:
		S()
	}
}
