package src

import (
	"fmt"
	"strconv"
)

// objects
const (
	VOID_O = iota
	NUM_O
	NAME_O
	STRING_O
	LAMBDA_O
	LIST_O
)

type Object struct {
	Type    int
	Content interface{}
	x       int
	y       int
}

func (t Object) ToStr() string {
	str := ""
	switch t.Type {
	case VOID_O:
		str += "type = VOID_O, "
	case NUM_O:
		str += "type = NUM_O, "
	case NAME_O:
		str += "type = NAME_O, "
	case STRING_O:
		str += "type = STRING_O, "
	case LAMBDA_O:
		str += "type = LAMBDA_O, "
	case LIST_O:
		str += "type = LIST_O, "
	}
	str += fmt.Sprintf("content = %v, ", t.Content)
	str += "x = " + strconv.Itoa(t.x) + ", "
	str += "y = " + strconv.Itoa(t.y) + "; "
	return str
}

type lambda struct {
	params []string
	body   interface{}
}
