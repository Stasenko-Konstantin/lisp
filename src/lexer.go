package src

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	x int
	y int
)

func Scan(code string) []Token {
	var (
		tokens  []Token
		symbols = "\\|/?.><!#@`^~%&*-_+=;"
		i       int
		x       = 0
		y       = 0
	)

	take := func(token, n, n2 int, p, pErr func(c rune) bool, err error) {
		content := ""
		for n, c := range code[i+n2:] {
			switch {
			case p(c):
				goto cont
			case pErr(c):
				lexerErr(err)
				if Repl {
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
			x:       x,
			y:       y,
		})
		x += len(content) + n - 1
	}

	for i < len(code) {
		switch c := code[i]; c {
		case '(', '[':
			tokens = append(tokens, Token{
				Type:    LPAREN_T,
				Content: "(",
				x:       x,
				y:       y,
			})
		case ')', ']':
			tokens = append(tokens, Token{
				Type:    RPAREN_T,
				Content: ")",
				x:       x,
				y:       y,
			})
		case '"':
			take(STRING_T, 2, 1,
				func(c rune) bool { return c == '"' },
				func(c rune) bool { return c == '\n' },
				errors.New("dangling \""))
		case '-':
			if code[i+1] == '-' {
				for {
					if code[i+1] == '\n' {
						break
					}
					i++
				}
			} else {
				take(NAME_T, 0, 0,
					func(c rune) bool { return !(isLetter(c) || strings.Contains(symbols, string(c)) || unicode.IsDigit(c)) },
					func(c rune) bool { return false }, nil)
			}
		case '\r', '\t', ' ':
			{
			}
		case '\n':
			x = -1
			y += 1
		default:
			switch {
			case isLetter(rune(c)) || strings.Contains(symbols, string(c)):
				take(NAME_T, 0, 0,
					func(c rune) bool { return !(isLetter(c) || strings.Contains(symbols, string(c)) || unicode.IsDigit(c)) },
					func(c rune) bool { return false }, nil)
			case unicode.IsDigit(rune(c)):
				take(NUM_T, 0, 0,
					func(c rune) bool { return !unicode.IsDigit(c) },
					func(c rune) bool { return false }, nil)
			default:
				lexerErr(errors.New("unknown symbol \"" + string(c) + "\""))
			}
		}
		x++
		i++
	}
	return tokens
}

func lexerErr(err error) {
	AddErr(errors.New(fmt.Sprintf("lexer error: %v; x = %d, y = %d\n", err, x, y)))
}
