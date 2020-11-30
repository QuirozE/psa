package h

import (
	"bufio"
	"io"
)

func NextToken(r *bufio.Reader) Token {
	c, _, err := r.ReadRune()
	if err == nil {
		switch c {
		case 'u':
			return U
		case 'v':
			return V
		case 'x':
			return X
		case 'y':
			return Y
		case 'z':
			return Z
		}
	} else if err == io.EOF {
		return Eof
	}
	return Err
}
