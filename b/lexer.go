package b

import (
	"bufio"
	"io"
)

func NextToken(r *bufio.Reader) Token {
	c, _, err := r.ReadRune()
	if err == nil {
		switch c {
		case '0':
			return Zero
		case '1':
			return One
		}
	} else if err == io.EOF {
		return Eof
	}
	return Err
}
