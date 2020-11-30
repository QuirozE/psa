package f

import (
	"bufio"
	"fmt"
)

var r *bufio.Reader
var currToken Token

func willAccept() {
	if r := recover(); r != nil {
		fmt.Println("Reject", r)
	} else if currToken != Eof {
		fmt.Println("Reject", currToken)
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
	case If:
		currToken = NextToken(r)
		E()
		S()
		T()
	case Otro:
		currToken = NextToken(r)
	default:
		panic("Syntax error S")
	}
}

func T() {
	switch currToken {
	case Else:
		currToken = NextToken(r)
		S()
	}
}

func E() {
	switch currToken {
	case True, False:
		currToken = NextToken(r)
	default:
		panic("Syntax error E")
	}
}
