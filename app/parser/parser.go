package parser

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/interpreter-starter-go/app/loxerrors"
	//lint:ignore ST1001 I dont care
	. "github.com/codecrafters-io/interpreter-starter-go/app/token"
)

// State of parser
var global_tokens []Token
var current_token int

func peek() Token {
	return global_tokens[current_token]
}

func advance() Token {
	if global_tokens[current_token].Type != EOF {
		current_token++
	}
	return previous()
}

func previous() Token {
	return global_tokens[current_token-1]
}

func match(Types ...TokenType) bool {
	for _, Type := range Types {
		if peek().Type == EOF {
			break
		}
		if peek().Type == Type {
			advance()
			return true
		}
	}

	return false
}

func consume(Type TokenType, message string) (Token, error) {
	if Type != EOF && peek().Type == Type {
		return advance(), nil
	}

	loxerrors.Parser_error(peek(), message)
	return Token{}, fmt.Errorf("parser error: %s", message)

}

// Synchronize after entering panic mode
func synchronize() {
	advance()

	for peek().Type != EOF {
		if previous().Type == SEMICOLON {
			return
		}

		switch peek().Type {
		case CLASS:
		case FUN:
		case VAR:
		case FOR:
		case IF:
		case WHILE:
		case PRINT:
		case RETURN:
			return
		}

		advance()
	}

}

// Parse given list of tokens and returns AST. Entry point to parser package.
func Parse(tokens []Token) (syntax_tree *AstNode, err error) {
	global_tokens = tokens

	root, err := expression()
	if err != nil {
		return nil, err
	}

	return root, nil
}

func AstPrinter(tree *AstNode) string {
	var builder strings.Builder

	switch tree.Type {
	case BINARY:
		builder.WriteString("(")

		val, ok := tree.Representation.(Token)
		if !ok {
			fmt.Println("Wrong type: Not Token when BINARY")
		}
		builder.WriteString(val.Lexeme + " ")

		left := AstPrinter(tree.Children[0])
		builder.WriteString(left + " ")

		right := AstPrinter(tree.Children[1])
		builder.WriteString(right)
		builder.WriteString(")")

	case UNARY:
		builder.WriteString("(")

		val, ok := tree.Representation.(Token)
		if !ok {
			fmt.Println("Wrong type: Not Token when UNARY")
		}
		builder.WriteString(val.Lexeme + " ")

		right := AstPrinter(tree.Children[0])
		builder.WriteString(right)
		builder.WriteString(")")

	case TERMINAL:
		val, ok := tree.Representation.(Token)
		if !ok {
			fmt.Println("Wrong type: Not Token when TERMINAL")
		}
		builder.WriteString(val.Lexeme)
	case LITERAL:
		val, ok := tree.Representation.(string)
		if !ok {
			fmt.Println("Wrong type: Not string when LITERAL")
		}
		builder.WriteString(val)

	case GROUP:
		builder.WriteString("(")
		builder.WriteString("group ")

		child := AstPrinter(tree.Children[0])
		builder.WriteString(child)
		builder.WriteString(")")

	}

	return builder.String()
}
