package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	parse, _ := Parse(Scan("(+ 1 2)"))
	assert.Equal(t, &Object{
		Type: LIST_O,
		Content: Program{
			{
				Type:    NAME_O,
				Content: interface{}("+"),
				x:       1,
			},
			{
				Type:    NUM_O,
				Content: interface{}(1),
				x:       3,
			},
			{
				Type:    NUM_O,
				Content: interface{}(2),
				x:       5,
			},
		},
	}, parse)
}
