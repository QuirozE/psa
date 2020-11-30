package d

type Token int

const (
	Id = iota
	Num
	Lpar
	Rpar
	Comma
	Eof
	Err
)
