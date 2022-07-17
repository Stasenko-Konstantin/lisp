package src

import (
	"errors"
	"fmt"
)

type Env struct {
	Parent *Env
	Defs   map[string]*Object
}

func (env Env) find(name string) *Object {
	if _, ok := env.Defs[name]; ok {
		return env.Defs[name]
	} else if env.Parent == nil {
		evalErr(MakeVoid(nil), errors.New("not found: "+name))
		return MakeVoid(nil)
	} else {
		return env.Parent.find(name)
	}
}

func Eval(obj *Object, env Env) *Object {
	switch obj.Type {
	case LAMBDA_O:
		return &Object{
			Type: VOID_O,
			x:    obj.x,
			y:    obj.y,
		}
	case NAME_O:
		return evalName(obj, env)
	case LIST_O:
		return evalList(obj, env)
	case BUILTIN_O:
		return evalBuiltin(obj, env)
	default:
		return obj // NUM_O, STRING_O and VOID_O
	}
}

func evalName(obj *Object, env Env) *Object {
	if v, ok := env.Defs[obj.Content.(string)]; ok {
		return v
	} else {
		if env.Parent == nil {
			evalErr(obj, errors.New("not found in this scope { "+obj.Content.(string)+" }"))
		} else {
			return evalName(obj, *env.Parent)
		}
	}
	return nil
}

func evalList(obj *Object, env Env) *Object {
	content := obj.Content.(Program)
	if len(content) == 0 {
		return MakeVoid(obj)
	}
	head := content[0]
	switch head.Type {
	case NAME_O:
		switch head.Content.(string) {
		case "define":
			return evalDefine(obj, env)
		case "lambda":
			return evalLambda(obj)
		default:
			return evalFunctionCall(obj, env, head.Content.(string))
		}
	default:
		var newList Program
		for _, o := range content {
			result := Eval(&o, env)
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
			x:       obj.x,
			y:       obj.y,
		}
	}
}

func evalDefine(obj *Object, env Env) *Object {
	var (
		name    = &Object{}
		content = obj.Content.(Program)
	)
	if len(content) < 3 {
		evalErr(obj, errors.New("invalid number of arguments for define"))
		return MakeVoid(obj)
	}
	switch content[1].Type {
	case NAME_O:
		name.Content = content[1].Content
	default:
		evalErr(obj, errors.New("invalid define"))
		return nil
	}
	val := Eval(&content[2], env)
	env.Defs[name.Content.(string)] = val
	return MakeVoid(obj)
}

func evalLambda(obj *Object) *Object {
	var (
		params  []string
		body    interface{}
		content = obj.Content.(Program)
	)
	switch content[1].Type {
	case LIST_O:
		ps := content[1].Content.(Program)
		for _, p := range ps {
			switch p.Type {
			case NAME_O:
				params = append(params, p.Content.(string))
			default:
				evalErr(obj, errors.New("invalid lambda parameter"))
				return MakeVoid(obj)
			}
		}
	default:
		evalErr(obj, errors.New("invalid lambda parameter"))
		return MakeVoid(obj)
	}
	switch content[2].Type {
	case LIST_O:
		body = content[2].Content
	default:
		evalErr(obj, errors.New("invalid lambda"))
		return MakeVoid(obj)
	}
	return &Object{
		Type: LAMBDA_O,
		Content: interface{}(lambda{
			params: params,
			body:   body,
		}),
		x: obj.x,
		y: obj.y,
	}
}

func evalFunctionCall(obj *Object, env Env, name string) *Object {
	content := obj.Content.(Program)
	value := env.find(name)
	switch value.Type {
	case LAMBDA_O:
		defs := make(map[string]*Object)
		lambda := value.Content.(lambda)
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
			x:       obj.x,
			y:       obj.y,
		}, newEnv)
	case BUILTIN_O:
		return evalBuiltin(obj, env)
	default:
		evalErr(obj, errors.New("not a lambda { "+name+" }"))
	}
	return MakeVoid(obj)
}

func evalBuiltin(obj *Object, env Env) *Object {
	content := obj.Content.(Program)
	f := env.find(content[0].Content.(string)).
		Content.(func(*Object, Env) *Object)
	for i, o := range content {
		switch o.Type {
		case NAME_O:
			content[i] = *env.find(o.Content.(string))
		}
	}
	return f(
		&Object{
			Type:    LIST_O,
			Content: content[1:],
			x:       0,
			y:       0,
		}, env)
}

func MakeVoid(obj *Object) *Object {
	return &Object{
		Type:    VOID_O,
		Content: obj,
		x:       0,
		y:       0,
	}
}

func evalErr(object *Object, err error) {
	AddErr(errors.New(fmt.Sprintf("runtime error: %v\n\tcontent = %v, x = %d, y = %d\n", err, object.Content, object.x, object.y)))
}
