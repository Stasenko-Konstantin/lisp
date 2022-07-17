package src

import (
	"fmt"
	"strconv"
	"strings"
)

// Type from Object struct
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

// GetContent - pretty printer for evaluated objects
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
		if t.Content.(*Object) == nil {
			return ""
		}
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
	return t.ToStr("")
}

// ToStr - pretty printer for ast
func (t Object) ToStr(tab string) string {
	str := tab
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
	var content interface{}
	switch t.Content.(type) {
	case string, int:
		content = t.Content
	case Program:
		var r string
		for _, o := range t.Content.(Program) {
			r += o.ToStr(tab + "\t")
		}
		content = r
	default:
		content = t.Content.(*Object).ToStr("\t")
	}
	str += fmt.Sprintf("content = %v, ", content)
	str += "x = " + strconv.Itoa(t.x) + ", "
	str += "y = " + strconv.Itoa(t.y) + ";\n"
	var r string
	for _, l := range strings.Split(str, "\n") {
		if l != "" {
			r += "\n" + l
		}
	}
	return r
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
			return MakeVoid(obj)
		},
		x: 0,
		y: 0,
	}
	defs["print"] = &Object{
		Type: BUILTIN_O,
		Content: func(obj *Object, env Env) *Object {
			fmt.Print(Eval(obj, env).GetContent(false))
			return MakeVoid(obj)
		},
		x: 0,
		y: 0,
	}
	defs["^"] = &Object{ // return
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
