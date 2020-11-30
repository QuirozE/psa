package h

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
	case U:
		currToken = NextToken(r)
		B()
		D()
		switch currToken {
		case Z:
			currToken = NextToken(r)
		default:
			panic("Syntax error")
		}
	default:
		panic("Syntax error")
	}
}

func B() {
	switch currToken {
	case V:
		currToken = NextToken(r)
		B()
	}
}

func D() {
	switch currToken {
	case Y:
		currToken = NextToken(r)
		switch currToken {
		case X:
			currToken = NextToken(r)
		}
	case X:
		currToken = NextToken(r)
	}
}
