package token

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func StringifyToken(token Token) string {
	var literal string
	if token.Literal == nil {
		literal = "null"
	} else if value, ok := token.Literal.(string); ok {
		literal = value
	} else if value, ok := token.Literal.(float64); ok {
		literal = strconv.FormatFloat(value, 'f', -1, 64)
		if !strings.ContainsAny(literal, ".") {
			literal += ".0"
		}
	}
	return fmt.Sprintf("%s %s %s", token.Type, token.Lexeme, literal)
}
