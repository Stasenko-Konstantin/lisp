package src

import "strconv"

const (
	NUM = iota
	NAME
	STRING
	LPAREN
	RPAREN
)

type Token struct {
	Type    int
	Content string
	X       int
	Y       int
}

func (t Token) ToStr() string {
	str := ""
	switch t.Type {
	case NUM:
		str += "type = NUM, "
	case NAME:
		str += "type = NAME, "
	case STRING:
		str += "type = STRING, "
	case RPAREN:
		str += "type = RPAREN, "
	case LPAREN:
		str += "type = LPAREN, "
	}
	str += "content = " + t.Content + ", "
	str += "x = " + strconv.Itoa(t.X) + ", "
	str += "y = " + strconv.Itoa(t.Y) + ";\n"
	return str
}
