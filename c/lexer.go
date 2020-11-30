package c

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

func consumeDigits(r *bufio.Reader) {
	for {
		c, _, _ := r.ReadRune()
		if !unicode.IsDigit(c) {
			r.UnreadRune()
			return
		}
	}
}

func NextToken(r *bufio.Reader) Token {
	c, _, err := r.ReadRune()
	if err == nil {
		switch c {
		case '(':
			return Lpar
		case ')':
			return Rpar
		default:
			if unicode.IsDigit(c) {
				consumeDigits(r)
				return Num
			} else if unicode.IsLetter(c) {
				consumeLetters(r)
				return Id
			}
		}

	} else if err == io.EOF {
		return Eof
	}
	return Err
}
