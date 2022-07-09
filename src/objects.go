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
	BUILTIN_O
)

type Object struct {
	Type    int
	Content interface{}
	x       int
	y       int
}

func (t *Object) GetContent(list bool) string {
	switch t.Content.(type) {
	case Program:
		r := ""
		for _, o := range t.Content.(Program) {
			r += o.GetContent(true) + " "
		}
		if list {
			return "( " + r + ")"
		}
		return r
	case *Object:
		switch t.Content.(*Object).Type {
		case LIST_O:
			r := "( "
			for _, o := range t.Content.(*Object).Content.(Program) {
				r += o.GetContent(true) + " "
			}
			return r + ")"
		}
		return t.Content.(*Object).GetContent(false)
	case int:
		return strconv.Itoa(t.Content.(int))
	case string:
		return t.Content.(string)
	}
	return t.ToStr()
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
	case BUILTIN_O:
		str += "type = BUILTIN_O, "
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

func MakeBuiltins() map[string]*Object {
	defs := make(map[string]*Object)
	defs["println"] = &Object{
		Type: BUILTIN_O,
		Content: func(obj *Object, env Env) *Object {
			fmt.Println(Eval(obj, env).GetContent(false))
			return makeVoid(obj)
		},
		x: 0,
		y: 0,
	}
	defs["^"] = &Object{
		Type: BUILTIN_O,
		Content: func(obj *Object, env Env) *Object {
			r := Eval(obj, env)
			return r
		},
		x: 0,
		y: 0,
	}
	return defs
}
