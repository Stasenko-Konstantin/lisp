package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := len(os.Args)
	if args == 1 {
		repl()
	} else if args == 2 {
		file := os.Args[1]
		if _, err := os.OpenFile(file, os.O_RDONLY, 0755); err == nil {
			if filepath.Ext(file) != ".lisp" {
				fileErr(errors.New("its not *.lisp\n"))
			}
		} else {
			fmt.Fprintf(os.Stderr, "cant read file\n"+err.Error())
		}
		data, err := os.ReadFile(file)
		fileErr(err)
		eval(string(data), false)
	} else {
		fmt.Fprintf(os.Stderr, "usage: ./lisp [*.lisp]\n")
	}
}

func fileErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "cant read file\n"+err.Error())
		os.Exit(1)
	}
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("< ")
		code, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("> " + err.Error())
		}
		code = strings.Replace(code, "\n", "", -1)
		fmt.Println(">", eval(code, true))
	}
}

func eval(code string, repl bool) interface{} {
	return code
}
