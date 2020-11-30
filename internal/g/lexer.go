package g

import (
	"bufio"
	"io"
	"unicode"
)

func consumeLetters(r *bufio.Reader) {
	for {
		c, _, _ := r.ReadRune()
		if !unicode.IsLetter(c) {
			r.UnreadRune()
			return
		}
	}
}

func NextToken(r *bufio.Reader) Token {
	c, _, err := r.ReadRune()
	if err == nil {
		switch c {
		case ':':
			return Ddot
		default:
			if unicode.IsLetter(c) {
				consumeLetters(r)
				return Id
			} else if unicode.IsSpace(c) {
				return NextToken(r)
			}
		}
	} else if err == io.EOF {
		return Eof
	}
	return Err
}
