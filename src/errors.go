package src

import (
	"fmt"
	"os"
)

var (
	Errors              []error
	InterpretationFault bool
	Repl                bool
)

func AddErr(err error) {
	InterpretationFault = true
	Errors = append(Errors, err)
}

func PrintErrors() {
	fmt.Println("Interpretation Fault:")
	for n, e := range Errors {
		fmt.Fprintf(os.Stderr, "\t%d, %s", n, e.Error())
	}
	if !Repl {
		os.Exit(1)
	}
}

func ResetErrors() {
	InterpretationFault = false
	Errors = []error{}
}
