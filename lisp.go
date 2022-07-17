package main

import (
	"bufio"
	"errors"
	"fmt"
	"lisp/src"
	"os"
	"path/filepath"
	"strings"
)

var (
	out   = ">>> "
	in    = "<<< "
	usage = "usage: ./lisp [*.scm] [arg] where arg:\n" +
		"\t--help \t -- prints the help\n" +
		"\t--past \t -- prints intermediate ast"
	help = out + ":l, :load *.scm \t -- evaluate file\n" +
		out + ":h, :help \t\t -- prints the help\n" +
		out + ":p, :past \t\t -- prints intermediate ast\n" +
		out + ":q, :quit \t\t -- program exit"
)

var (
	printAst bool
	env      src.Env
)

func main() {
	args := os.Args
	sargs := strings.Join(args, " ")
	defs := src.MakeBuiltins()
	env = src.Env{
		Parent: nil,
		Defs:   defs,
	}
	if len(args) == 1 {
		repl()
	} else {
		if strings.Contains(sargs, "--help") {
			fmt.Println(usage)
			os.Exit(0)
		} else if strings.Contains(sargs, "--past") {
			printAst = true
		}
		src.Repl = false
		evalFile(args[1])
	}
}

func evalFile(file string) error {
	if _, err := os.OpenFile(file, os.O_RDONLY, 0755); err == nil {
		if filepath.Ext(file) != ".scm" {
			if err := fileErr(errors.New("its not *.scm")); err != nil {
				return err
			}
		}
	} else {
		fmt.Fprintf(os.Stderr, "cant read file\n%s\n", err.Error())
	}
	data, err := os.ReadFile(file)
	fileErr(err)
	eval("("+string(data)+")", true)
	return nil
}

func fileErr(err error) error {
	if err != nil {
		if src.Repl {
			src.AddErr(errors.New(fmt.Sprintf("cant read file:\n\t%s\n", err.Error())))
			return err
		} else {
			fmt.Fprintf(os.Stderr, "cant read file:\n%s\n", err.Error())
			os.Exit(1)
		}
	}
	return err
}

func repl() {
	src.Repl = true
	reader := bufio.NewReader(os.Stdin)
	for {
		src.ResetErrors()
		fmt.Print(in)
		code, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(out + err.Error())
		}
		codes := strings.Split(code[:len(code)-1], " ")
		switch codes[0] {
		case ":h", ":help":
			fmt.Println(help)
			continue
		case ":p", ":past":
			printAst = !printAst
			if printAst {
				fmt.Println(out + "OK")
			} else {
				fmt.Println(out + "NO")
			}
			continue
		case ":l", ":load":
			for _, f := range codes[1:] {
				if f != "" {
					err := evalFile(f)
					if err != nil {
						fmt.Println(strings.TrimSuffix(out, " "), eval("", false))
					}
					break
				}
			}
			continue
		case ":q", ":quit":
			os.Exit(0)
		default:
			fmt.Println(strings.TrimSuffix(out, " "), eval(code, true))
		}
	}
}

func eval(code string, needEval bool) interface{} {
	tokens := src.Scan(code + " ")
	var objects *src.Object
	if tokens == nil {
		objects = src.MakeVoid(nil)
	} else {
		objects, _ = src.Parse(tokens)
	}
	if printAst {
		fmt.Println(objects.ToStr("") + "\n")
	}
	var obj *src.Object
	if needEval {
		obj = src.Eval(objects, env)
	} else {
		obj = src.MakeVoid(nil)
	}
	if src.InterpretationFault {
		src.PrintErrors()
	}
	return obj.GetContent(false)
}
