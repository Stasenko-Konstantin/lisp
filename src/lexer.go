package src

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var (
	x int
	y int
)

func Scan(code string, repl bool) []Token {
	var (
		tokens  []Token
		symbols = "\\|/?.><!#@`^~%&*-_+=;"
		i       int
	)

	take := func(token, n, n2 int, p, pErr func(c rune) bool, err error) {
		content := ""
		for n, c := range code[i+n2:] {
			switch {
			case p(c):
				goto cont
			case pErr(c):
				lexerErr(err, repl)
				if repl {
					code = code[n+i+n2:]
					i += 1
					continue
				}
			default:
				content += string(c)
			}
		}
	cont:
		code = code[len(content)+i+n:]
		i = -1
		tokens = append(tokens, Token{
			Type:    token,
			Content: content,
			X:       x,
			Y:       y,
		})
		x += len(content) + n
	}

	x = 0
	for i < len(code) {
		switch c := code[i]; c {
		case '(', '[':
			tokens = append(tokens, Token{
				Type:    LPAREN,
				Content: "(",
				X:       x,
				Y:       y,
			})
		case ')', ']':
			tokens = append(tokens, Token{
				Type:    RPAREN,
				Content: ")",
				X:       x,
				Y:       y,
			})
		case '"':
			take(STRING, 2, 1,
				func(c rune) bool { return c == '"' },
				func(c rune) bool { return c == '\n' },
				errors.New("dangling \""))
		case '\r', '\t', ' ':
			{
			}
		case '\n':
			x = -1
			y += 1
		default:
			switch {
			case unicode.IsLetter(rune(c)) || strings.Contains(symbols, string(c)):
				take(NAME, 0, 0,
					func(c rune) bool { return !(unicode.IsLetter(c) || strings.Contains(symbols, string(c))) },
					func(c rune) bool { return false }, nil)
			case unicode.IsDigit(rune(c)):
				take(NUM, 0, 0,
					func(c rune) bool { return !unicode.IsDigit(c) },
					func(c rune) bool { return false }, nil)
			default:
				lexerErr(errors.New("unknown symbol \""+string(c)+"\""), repl)
			}
		}
		x++
		i++
	}
	return tokens
}

func lexerErr(err error, repl bool) {
	fmt.Fprintf(os.Stderr, "> lexer error: %v\n> x = %d, y = %d\n", err, x, y)
	if !repl {
		os.Exit(1)
	}
}
