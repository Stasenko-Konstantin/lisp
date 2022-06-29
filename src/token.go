package src

const (
	NUM = iota
	NAME
	STRING
	LPARENT
	RPARENT
)

type Token struct {
	Type    int
	Content string
	X       int
	Y       int
}
