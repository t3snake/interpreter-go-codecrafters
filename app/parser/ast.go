package parser

import (
	"errors"

	"github.com/codecrafters-io/interpreter-starter-go/app/loxerrors"
	//lint:ignore ST1001 I dont care
	. "github.com/codecrafters-io/interpreter-starter-go/app/token"
)

type NodeType string

const (
	BINARY   NodeType = "binary"
	UNARY    NodeType = "unary"
	TERMINAL NodeType = "terminal"
	LITERAL  NodeType = "literal"
	GROUP    NodeType = "group"
)

// Abstract Syntax Tree Node
type AstNode struct {
	Representation any
	Type           NodeType
	Children       []*AstNode
}

/*  Context Free Grammer from low to high precedence:

expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil"
               | "(" expression ")" ;
*/

func expression() (*AstNode, error) {
	return equality()
}

func equality() (*AstNode, error) {
	var expr *AstNode
	expr, err := comparison()
	if err != nil {
		return nil, err
	}

	for match(BANG_EQUAL, EQUAL_EQUAL) {
		right, err := comparison()
		if err != nil {
			return nil, err
		}

		expr = &AstNode{
			Representation: previous(),
			Type:           BINARY,
			Children:       []*AstNode{expr, right},
		}
	}

	return expr, nil
}

func comparison() (*AstNode, error) {
	var expr *AstNode
	expr, err := term()
	if err != nil {
		return nil, err
	}

	for match(LESS, LESS_EQUAL, GREATER, GREATER_EQUAL) {
		right, err := term()
		if err != nil {
			return nil, err
		}

		expr = &AstNode{
			Representation: previous(),
			Type:           BINARY,
			Children:       []*AstNode{expr, right},
		}
	}

	return expr, nil
}

func term() (*AstNode, error) {
	var expr *AstNode
	expr, err := factor()
	if err != nil {
		return nil, err
	}

	for match(MINUS, PLUS) {
		right, err := factor()
		if err != nil {
			return nil, err
		}

		expr = &AstNode{
			Representation: previous(),
			Type:           BINARY,
			Children:       []*AstNode{expr, right},
		}
	}

	return expr, nil
}

func factor() (*AstNode, error) {
	var expr *AstNode
	expr, err := unary()
	if err != nil {
		return nil, err
	}

	for match(STAR, SLASH) {
		right, err := unary()
		if err != nil {
			return nil, err
		}

		expr = &AstNode{
			Representation: previous(),
			Type:           BINARY,
			Children:       []*AstNode{expr, right},
		}
	}

	return expr, nil
}

func unary() (*AstNode, error) {
	if match(BANG, MINUS) {
		child, err := unary()
		if err != nil {
			return nil, err
		}

		expr := &AstNode{
			Representation: previous(),
			Type:           UNARY,
			Children:       []*AstNode{child},
		}

		return expr, nil
	}

	return primary()
}

func primary() (*AstNode, error) {
	if match(TRUE) {
		return &AstNode{
			Representation: previous(),
			Type:           TERMINAL,
			Children:       nil,
		}, nil
	}

	if match(FALSE) {
		return &AstNode{
			Representation: previous(),
			Type:           TERMINAL,
			Children:       nil,
		}, nil
	}

	if match(NIL) {
		return &AstNode{
			Representation: previous(),
			Type:           TERMINAL,
			Children:       nil,
		}, nil
	}

	if match(IDENTIFIER) || match(NUMBER) || match(STRING) {
		return &AstNode{
			Representation: previous().Literal,
			Type:           TERMINAL,
			Children:       nil,
		}, nil
	}

	if match(LEFT_PAREN) {
		expr, err := expression()
		if err != nil {
			return nil, err
		}

		expr = &AstNode{
			Representation: "group",
			Type:           GROUP,
			Children:       []*AstNode{expr},
		}

		_, err = consume(RIGHT_PAREN, "Expected ')' after expression.")

		return expr, err
	}

	loxerrors.Parser_error(peek(), "Expect expression.")
	return nil, errors.New(`Unidentified token: ` + peek().Lexeme)

}
