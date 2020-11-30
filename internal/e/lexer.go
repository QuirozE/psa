package e

import (
	"bufio"
	"io"
)

func NextToken(r *bufio.Reader) Token {
	c, _, err := r.ReadRune()
	if err == nil {
		switch c {
		case 'b':
			return B
		case 'c':
			return C
		case 'f':
			return F
		case 'a':
			return A
		}
	} else if err == io.EOF {
		return Eof
	}
	return Err
}
