package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	parse, _ := Parse(Scan("(1 2)"))
	eval := Eval(parse, Env{
		Parent: nil,
		Defs:   MakeBuiltins(),
	})
	assert.Equal(t, &Object{
		Type: LIST_O,
		Content: Program{
			Object{
				Type:    NUM_O,
				Content: 1,
				x:       1,
				y:       0,
			},
			Object{
				Type:    NUM_O,
				Content: 2,
				x:       4,
				y:       0,
			},
		},
		x: 0,
		y: 0,
	}, eval)
}
