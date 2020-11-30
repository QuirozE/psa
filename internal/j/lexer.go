package j

import (
	"bufio"
	"io"
)

func NextToken(r *bufio.Reader) Token {
	c, _, err := r.ReadRune()
	if err == nil {
		switch c {
		case 'a':
			return A
		case 'b':
			return B
		case '(':
			return Lpar
		case ')':
			return Rpar
		case '*':
			return Star
		case '?':
			return Option
		case '%':
			return Mod
		}
	} else if err == io.EOF {
		return Eof
	}
	return Err
}
