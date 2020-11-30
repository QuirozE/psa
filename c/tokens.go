package c

type Token int

const (
	Id = iota
	Num
	Lpar
	Rpar
	Eof
	Err
)
