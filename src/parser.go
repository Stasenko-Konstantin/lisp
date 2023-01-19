package src

import (
	"errors"
	"fmt"
	"strconv"
)

type Program []Object

type parser struct {
	program Program
	tokens  []Token
	skip    int
}

// skipnt - skip some tokens count
func Parse(tokens []Token) (o *Object, skipnt int) {
	p := parser{
		tokens: tokens,
		skip:   -1,
	}
	return p.parse()
}

func (p *parser) add(object ObjectType, content interface{}, token Token) {
	p.program = append(p.program, Object{
		Type:    object,
		Content: content,
		x:       token.x,
		y:       token.y,
	})
}

// skipnt - skip some tokens count
func (p *parser) parse() (o *Object, skipnt int) {
	if len(p.tokens) == 0 {
		return nil, 0
	}
	if !p.tokens[0].compare(Token{
		Type:    LPAREN_T,
		Content: "(",
	}) {
		parserErr(p.tokens[0], errors.New("should be ("))
	}
	for n, t := range p.tokens[1:] {
		if n <= p.skip {
			continue
		}
		switch t.Type {
		case NUM_T:
			content, _ := strconv.Atoi(t.Content)
			p.add(NUM_O, content, t)
		case NAME_T:
			p.add(NAME_O, t.Content, t)
		case STRING_T:
			p.add(STRING_O, t.Content, t)
		case LPAREN_T:
			subList, skipnt := Parse(p.tokens[n+1:])
			p.skip = skipnt + n + 1
			p.program = append(p.program, *subList)
		case RPAREN_T:
			return &Object{
				Type:    LIST_O,
				Content: p.program,
			}, n
		}
	}
	return &Object{
		Type:    LIST_O,
		Content: p.program,
	}, 0
}

func parserErr(token Token, err error) {
	AddErr(errors.New(fmt.Sprintf("parser error: %v\n\tcontent = %s, x = %d, y = %d\n", err, token.Content, token.x, token.y)))
}
