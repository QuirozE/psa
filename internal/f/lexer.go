package f

import (
	"bufio"
	"io"
	"strings"
)

func NextToken(r *bufio.Reader) Token {
	c, err := r.ReadString(' ')
	c = strings.Trim(c, " ")
	if c == "" && err == io.EOF {
		return Eof
	}

	if err == nil || err == io.EOF {
		switch c {
		case "if":
			return If
		case "else":
			return Else
		case "otro":
			return Otro
		case "true":
			return True
		case "false":
			return False
		}
	}
	return Err
}
