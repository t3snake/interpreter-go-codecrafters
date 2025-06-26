package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	token_type TokenType
	lexeme     string
	literal    any
	line       int
}

func stringifyToken(token Token) string {
	var literal string
	if token.literal == nil {
		literal = "null"
	} else if value, ok := token.literal.(string); ok {
		literal = value
	} else if value, ok := token.literal.(float64); ok {
		literal = strconv.FormatFloat(value, 'f', -1, 64)
		if !strings.ContainsAny(literal, ".") {
			literal += ".0"
		}
	}
	return fmt.Sprintf("%s %s %s", token.token_type, token.lexeme, literal)
}
