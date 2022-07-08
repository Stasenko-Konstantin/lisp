package main

import (
	"bufio"
	"errors"
	"fmt"
	"lisp/src"
	"os"
	"path/filepath"
)

var (
	usage = "usage: ./lisp [*.lisp]"
)

var env src.Env

func main() {
	args := len(os.Args)
	defs := src.MakeBuiltins()
	env = src.Env{
		Parent: nil,
		Defs:   defs,
	}
	if args == 1 {
		repl()
	} else if args == 2 {
		if os.Args[1] == "--help" {
			fmt.Println(usage)
			os.Exit(0)
		}
		file := os.Args[1]
		if _, err := os.OpenFile(file, os.O_RDONLY, 0755); err == nil {
			if filepath.Ext(file) != ".lisp" {
				fileErr(errors.New("its not *.lisp"))
			}
		} else {
			fmt.Fprintf(os.Stderr, "cant read file\n%s\n", err.Error())
		}
		data, err := os.ReadFile(file)
		fileErr(err)
		eval("("+string(data)+")", false)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", usage)
	}
}

func fileErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "cant read file:\n%s\n", err.Error())
		os.Exit(1)
	}
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		src.ResetErrors()
		fmt.Print("< ")
		code, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("> " + err.Error())
		}
		fmt.Println(">", eval(code, true))
	}
}

func eval(code string, repl bool) interface{} {
	src.Repl = repl
	tokens := src.Scan(code + " ")
	//fmt.Println(func(tokens []src.Token) []string {
	//	var r []string
	//	for _, t := range tokens {
	//		r = append(r, t.ToStr())
	//	}
	//	return r
	//}(tokens))
	objects, _ := src.Parse(tokens)
	obj := src.Eval(objects, env)
	if src.InterpretationFault {
		src.PrintErrors()
	}
	return obj.GetContent(false)
}
