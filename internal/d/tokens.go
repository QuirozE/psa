package d

type Token int

const (
	Id = iota
	Num
	Comma
	Lpar
	Rpar
	Eof
	Err
)
