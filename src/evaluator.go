package src

import (
	"errors"
	"fmt"
)

type Env struct {
	Parent *Env
	Defs   map[string]*Object
}

func Eval(object *Object, env Env) *Object {
	switch object.Type {
	case LAMBDA_O:
		return &Object{
			Type: VOID_O,
			x:    object.x,
			y:    object.y,
		}
	case NAME_O:
		return evalName(object, env)
	case LIST_O:
		return evalList(object, env)
	default:
		return object // NUM_O, STRING_O and VOID_O
	}
}

func evalName(name *Object, env Env) *Object {
	if v, ok := env.Defs[name.Content.(string)]; ok {
		return v
	} else {
		if env.Parent == nil {
			evalErr(name, errors.New("not found in this scope { "+name.Content.(string)+" }"))
		} else {
			return evalName(name, *env.Parent)
		}
	}
	return nil
}

func evalList(list *Object, env Env) *Object {
	content := list.Content.(Program)
	if len(content) == 0 {
		return makeVoid(list)
	}
	head := content[0]
	switch head.Type {
	case NAME_O:
		switch head.Content.(string) {
		case "define":
			return evalDefine(list, env)
		case "lambda":
			return evalLambda(list)
		default:
			return evalFunctionCall(list, env, head.Content.(string))
		}
	default:
		var newList Program
		for _, object := range content {
			result := Eval(&object, env)
			switch result.Type {
			case VOID_O:
				{
				}
			default:
				newList = append(newList, *result)
			}
		}
		return &Object{
			Type:    LIST_O,
			Content: interface{}(newList),
			x:       list.x,
			y:       list.y,
		}
	}
}

func evalDefine(list *Object, env Env) *Object {
	var (
		name    = &Object{}
		content = list.Content.(Program)
	)
	if len(content) < 3 {
		evalErr(list, errors.New("invalid number of arguments for define"))
		return makeVoid(list)
	}
	switch content[1].Type {
	case NAME_O:
		name.Content = content[1].Content
	default:
		evalErr(list, errors.New("invalid define"))
		return nil
	}
	val := Eval(&content[2], env)
	env.Defs[name.Content.(string)] = val
	return makeVoid(list)
}

func evalLambda(list *Object) *Object {
	var (
		params  []string
		body    interface{}
		content = list.Content.(Program)
	)
	switch content[1].Type {
	case LIST_O:
		ps := content[1].Content.(Program)
		for _, p := range ps {
			switch p.Type {
			case NAME_O:
				params = append(params, p.Content.(string))
			default:
				evalErr(list, errors.New("invalid lambda parameter"))
				return makeVoid(list)
			}
		}
	default:
		evalErr(list, errors.New("invalid lambda parameter"))
		return makeVoid(list)
	}
	switch content[2].Type {
	case LIST_O:
		body = content[2].Content
	default:
		evalErr(list, errors.New("invalid lambda"))
		return makeVoid(list)
	}
	return &Object{
		Type: LAMBDA_O,
		Content: interface{}(lambda{
			params: params,
			body:   body,
		}),
		x: list.x,
		y: list.y,
	}
}

func evalFunctionCall(list *Object, env Env, name string) *Object {
	content := list.Content.(Program)
	if v, ok := env.Defs[name]; ok {
		switch v.Type {
		case LAMBDA_O:
			defs := make(map[string]*Object)
			lambda := v.Content.(lambda)
			newEnv := Env{
				Parent: &env,
				Defs:   defs,
			}
			for i, p := range lambda.params {
				val := Eval(&content[i+1], env)
				newEnv.Defs[p] = val
			}
			return Eval(&Object{
				Type:    LIST_O,
				Content: lambda.body,
				x:       list.x,
				y:       list.y,
			}, newEnv)
		default:
			evalErr(list, errors.New("not a lambda { "+name+" }"))
			return makeVoid(list)
		}
	} else {
		if env.Parent == nil {
			evalErr(list, errors.New("not found in this scope { "+name+" }"))
		} else {
			return evalFunctionCall(list, *env.Parent, name)
		}
	}
	return makeVoid(list)
}

func makeVoid(list *Object) *Object {
	return makeVoid(list)
}

func evalErr(object *Object, err error) {
	addErr(errors.New(fmt.Sprintf("runtime error: %v\n\tcontent = %v, x = %d, y = %d\n", err, object.Content, object.x, object.y)))
}
