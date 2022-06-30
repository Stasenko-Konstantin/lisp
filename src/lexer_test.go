package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	assert.Equal(t, []Token{
		{
			Type:    NUM,
			Content: "123",
			X:       0,
			Y:       0,
		},
		{
			Type:    NAME,
			Content: "hello",
			X:       5,
			Y:       0,
		},
		{
			Type:    LPAREN,
			Content: "(",
			X:       12,
			Y:       0,
		},
		{
			Type:    RPAREN,
			Content: ")",
			X:       13,
			Y:       0,
		},
		{
			Type:    STRING,
			Content: "hello",
			X:       0,
			Y:       1,
		},
	}, Scan("123 hello (]\n\"hello\"", false))
}
