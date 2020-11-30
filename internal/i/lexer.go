package i

import (
	"bufio"
	"io"
)

func NextToken(r *bufio.Reader) Token {
	c, _, err := r.ReadRune()
	if err == nil {
		switch c {
		case '(':
			return Lpar
		case ')':
			return Rpar
		case ']':
			return Rbra
		}
	} else if err == io.EOF {
		return Eof
	}
	return Err
}
