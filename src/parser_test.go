package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
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
				x:       4,
			},
			{
				Type:    NUM_O,
				Content: interface{}(2),
				x:       7,
			},
		},
	}, Parse(Scan("(+ 1 2)")))
}
