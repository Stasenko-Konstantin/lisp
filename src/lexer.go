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

type lexer struct {
	tokens []Token
	code   string
	span   struct {
		x int
		y int
	}
	idx int
}

func Scan(code string) []Token {
	l := lexer{
		code: code,
		span: struct {
			x int
			y int
		}{0, 0},
	}
	return l.scan()
}

func (l *lexer) push(token int, content string) {
	l.tokens = append(l.tokens, Token{
		Type:    token,
		Content: content,
		x:       l.span.x,
		y:       l.span.y,
	})
}

func (l *lexer) take(token, n, n2 int, p, pErr func(c rune) bool, err error) {
	content := ""
	for n, c := range l.code[l.idx+n2:] {
		switch {
		case p(c):
			goto cont // break dont work here 🤔
		case pErr(c):
			lexerErr(err)
			if Repl {
				l.code = l.code[n+l.idx+n2:]
				l.idx += 1
				continue
			}
		default:
			content += string(c)
		}
	}
cont:
	l.code = l.code[len(content)+l.idx+n:]
	l.idx = -1
	l.push(token, content)
	l.span.x += len(content) + n - 1
}

func (l *lexer) scan() []Token {
	symbols := "\\|/?.><!#@`^~%&*-_+=;"

	for l.idx < len(l.code) {
		switch c := l.code[l.idx]; c {
		case '(', '[':
			l.push(LPAREN_T, "(")
		case ')', ']':
			l.push(RPAREN_T, ")")
		case '"':
			l.take(STRING_T, 2, 1,
				func(c rune) bool { return c == '"' },
				func(c rune) bool { return c == '\n' },
				errors.New("dangling \""))
		case '-':
			if l.code[l.idx+1] == '-' {
				for {
					if l.code[l.idx+1] == '\n' {
						break
					}
					l.idx++
				}
			} else {
				l.take(NAME_T, 0, 0,
					func(c rune) bool { return !(isLetter(c) || strings.Contains(symbols, string(c)) || unicode.IsDigit(c)) },
					func(c rune) bool { return false }, nil)
			}
		case '\r', '\t', ' ':
			{
			}
		case '\n':
			l.span.x = -1
			l.span.y += 1
		default:
			switch {
			case isLetter(rune(c)) || strings.Contains(symbols, string(c)):
				l.take(NAME_T, 0, 0,
					func(c rune) bool { return !(isLetter(c) || strings.Contains(symbols, string(c)) || unicode.IsDigit(c)) },
					func(c rune) bool { return false }, nil)
			case unicode.IsDigit(rune(c)):
				l.take(NUM_T, 0, 0,
					func(c rune) bool { return !unicode.IsDigit(c) },
					func(c rune) bool { return false }, nil)
			default:
				lexerErr(errors.New("unknown symbol \"" + string(c) + "\""))
			}
		}
		l.span.x++
		l.idx++
	}
	return l.tokens
}

func isLetter(r rune) bool {
	return (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
}

func lexerErr(err error) {
	AddErr(errors.New(fmt.Sprintf("lexer error: %v; x = %d, y = %d\n", err, x, y)))
}
