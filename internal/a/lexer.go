package a

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
		case '*':
			return Str
		case '+':
			return Pl
		}
	} else if err == io.EOF {
		return Eof
	}
	return Err
}
