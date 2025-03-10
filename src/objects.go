package src

import (
	"fmt"
	"strconv"
	"strings"
)

type ObjectType int

// Type from Object struct
const (
	VOID_O ObjectType = iota
	NUM_O
	NAME_O
	STRING_O
	LAMBDA_O
	LIST_O
	BUILTIN_O
)

type Object struct {
	Type    ObjectType
	Content interface{}
	x       int
	y       int
}

// GetContent - pretty printer for evaluated objects
func (o *Object) GetContent(list bool) string {
	switch o.Content.(type) {
	case Program:
		r := ""
		for _, o := range o.Content.(Program) {
			r += o.GetContent(true) + " "
		}
		if list {
			return "( " + r + ")"
		}
		return r
	case *Object:
		if o.Content.(*Object) == nil {
			return ""
		}
		switch o.Content.(*Object).Type {
		case LIST_O:
			r := "( "
			for _, o := range o.Content.(*Object).Content.(Program) {
				r += o.GetContent(true) + " "
			}
			return r + ")"
		}
		return o.Content.(*Object).GetContent(false)
	case int:
		return strconv.Itoa(o.Content.(int))
	case string:
		return o.Content.(string)
	}
	return o.AstStr("")
}

// AstStr - pretty printer for ast
func (o *Object) AstStr(tab string) string {
	str := tab
	switch o.Type {
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
	switch o.Content.(type) {
	case string, int:
		content = o.Content
	case Program:
		var r string
		for _, o := range o.Content.(Program) {
			r += o.AstStr(tab + "\t")
		}
		content = r
	default:
		content = o.Content.(*Object).AstStr("\t")
	}
	str += fmt.Sprintf("content = %v, ", content)
	str += "x = " + strconv.Itoa(o.x) + ", "
	str += "y = " + strconv.Itoa(o.y) + ";\n"
	var r string
	for _, l := range strings.Split(str, "\n") {
		if l != "" {
			r += "\n" + l
		}
	}
	return r
}

func (o *Object) String() string {
	return o.AstStr("")
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
	}
	defs["print"] = &Object{
		Type: BUILTIN_O,
		Content: func(obj *Object, env Env) *Object {
			fmt.Print(Eval(obj, env).GetContent(false))
			return MakeVoid(obj)
		},
	}
	defs["^"] = &Object{ // return
		Type: BUILTIN_O,
		Content: func(obj *Object, env Env) *Object {
			r := Eval(obj, env)
			return r
		},
	}
	return defs
}
