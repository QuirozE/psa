package j

type Token int

const (
	A = iota
	B
	Rpar
	Lpar
	Star
	Option
	Not
	Mod
	Eof
	Err
)
