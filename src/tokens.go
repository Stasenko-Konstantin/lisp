package src

import "strconv"

// tokens
const (
	NUM_T = iota
	NAME_T
	STRING_T
	LPAREN_T
	RPAREN_T
)

type Token struct {
	Type    int
	Content string
	x       int
	y       int
}

func (t Token) ToStr() string {
	str := ""
	switch t.Type {
	case NUM_T:
		str += "type = NUM_T, "
	case NAME_T:
		str += "type = NAME_T, "
	case STRING_T:
		str += "type = STRING_T, "
	case RPAREN_T:
		str += "type = RPAREN_T, "
	case LPAREN_T:
		str += "type = LPAREN_T, "
	}
	str += "content = " + t.Content + ", "
	str += "x = " + strconv.Itoa(t.x) + ", "
	str += "y = " + strconv.Itoa(t.y) + ";\n"
	return str
}

func (t Token) compare(that Token) bool {
	if t.Type != that.Type {
		return false
	} else if t.Content != that.Content {
		return false
	}
	return true
}
