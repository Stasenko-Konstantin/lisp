package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	assert.Equal(t, []Token{
		{
			Type:    NUM_T,
			Content: "123",
			x:       0,
			y:       0,
		},
		{
			Type:    NAME_T,
			Content: "hello",
			x:       5,
			y:       0,
		},
		{
			Type:    LPAREN_T,
			Content: "(",
			x:       12,
			y:       0,
		},
		{
			Type:    RPAREN_T,
			Content: ")",
			x:       13,
			y:       0,
		},
		{
			Type:    STRING_T,
			Content: "hello",
			x:       0,
			y:       1,
		},
	}, Scan("123 hello (]\n\"hello\""))
}
