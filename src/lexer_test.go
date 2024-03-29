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
		},
		{
			Type:    NAME_T,
			Content: "hello",
			x:       4,
		},
		{
			Type:    LPAREN_T,
			Content: "(",
			x:       10,
		},
		{
			Type:    RPAREN_T,
			Content: ")",
			x:       11,
		},
		{
			Type:    STRING_T,
			Content: "hello",
			y:       1,
		},
	}, Scan("123 hello (]\n\"hello\""))
}
