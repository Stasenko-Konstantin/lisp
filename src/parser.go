package src

import (
	"errors"
	"fmt"
	"strconv"
)

type Program []Object

// skipnt - skip some tokens count
func Parse(tokens []Token) (o *Object, skipnt int) {
	if len(tokens) == 0 {
		return nil, 0
	}
	var (
		program Program
		skip    = -1
	)
	if !tokens[0].compare(Token{
		Type:    LPAREN_T,
		Content: "(",
		x:       0,
		y:       0,
	}) {
		parserErr(tokens[0], errors.New("should be ("))
	}
	for n, t := range tokens[1:] {
		if n <= skip {
			continue
		}
		switch t.Type {
		case NUM_T:
			content, _ := strconv.Atoi(t.Content)
			program = append(program, Object{
				Type:    NUM_O,
				Content: content,
				x:       t.x,
				y:       t.y,
			})
		case NAME_T:
			program = append(program, Object{
				Type:    NAME_O,
				Content: t.Content,
				x:       t.x,
				y:       t.y,
			})
		case STRING_T:
			program = append(program, Object{
				Type:    STRING_O,
				Content: t.Content,
				x:       t.x,
				y:       t.y,
			})
		case LPAREN_T:
			subList, skipnt := Parse(tokens[n+1:])
			skip = skipnt + n + 1
			program = append(program, *subList)
		case RPAREN_T:
			return &Object{
				Type:    LIST_O,
				Content: program,
				x:       0,
				y:       0,
			}, n
		}
	}
	return &Object{
		Type:    LIST_O,
		Content: program,
		x:       0,
		y:       0,
	}, 0
}

func parserErr(token Token, err error) {
	AddErr(errors.New(fmt.Sprintf("parser error: %v\n\tcontent = %s, x = %d, y = %d\n", err, token.Content, token.x, token.y)))
}
