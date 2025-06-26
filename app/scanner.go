package main

import (
	"fmt"
)

type ScanState struct {
	start   int
	current int
	line    int
}

// globals
var source string
var scan_state ScanState
var tokens []Token

func scanTokens(input_source string) []Token {
	scan_state = ScanState{
		start:   0,
		current: 0,
		line:    1,
	}

	source = input_source
	tokens = make([]Token, 0)

	for !isAtEnd() {
		scan_state.start = scan_state.current
		scanToken()
	}
	return tokens
}

func scanToken() {
	char := advance()

	switch char {
	case '(':
		addToken(LEFT_PAREN)
	case ')':
		addToken(RIGHT_PAREN)
	case '{':
		addToken(LEFT_BRACE)
	case '}':
		addToken(RIGHT_BRACE)
	case ',':
		addToken(COMMA)
	case '.':
		addToken(DOT)
	case '-':
		addToken(MINUS)
	case '+':
		addToken(PLUS)
	case ';':
		addToken(SEMICOLON)
	case '*':
		addToken(STAR)

	// Characters that can be single char or resolve multiple
	case '!':
		if match('=') {
			addToken(BANG_EQUAL)
		} else {
			addToken(BANG)
		}
	case '=':
		if match('=') {
			addToken(EQUAL_EQUAL)
		} else {
			addToken(EQUAL)
		}
	case '<':
		if match('=') {
			addToken(LESS_EQUAL)
		} else {
			addToken(LESS)
		}
	case '>':
		if match('=') {
			addToken(GREATER_EQUAL)
		} else {
			addToken(GREATER)
		}

	default:
		error(scan_state.line, fmt.Sprintf("Unexpected character: %s", string(char)))
	}
}

// Return true if current pointer is at end
func isAtEnd() bool {
	return scan_state.current >= len(source)
}

// Return current rune and increment current pointer
func advance() rune {
	scan_state.current++
	return rune(source[scan_state.current-1])
}

// Return true if character at current pointer during scan is equal to the given rune
func match(next_char rune) bool {
	if scan_state.current >= len(source) {
		return false
	}
	return rune(source[scan_state.current]) == next_char
}

func addToken(token_type TokenType) {
	addTokenWithLiteral(token_type, nil)
}

func addTokenWithLiteral(token_type TokenType, literal any) {
	current_text := source[scan_state.start:scan_state.current]

	tokens = append(tokens, Token{
		token_type: token_type,
		lexeme:     current_text,
		literal:    literal,
		line:       scan_state.line,
	})
}
