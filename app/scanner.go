package main

import (
	"fmt"
	"strconv"
)

// State of current running scan
type ScanState struct {
	start   int // Start of token being scanned
	current int // Current unconsumed character
	line    int // Current line being scanned
}

// globals
var source string
var scan_state ScanState
var tokens []Token

// Scan and return list of tokens given source file
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

// Scan single token and adjust scan_state pointers
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
	case '/':
		if match('/') {
			// comment case
			for !isAtEnd() && peek() != '\n' {
				advance()
			}
		} else {
			addToken(SLASH)
		}

	// Whitespace handling
	case ' ':
		// Ignore
	case '\r':
		// Ignore
	case '\t':
		// Ignore
	case '\n':
		scan_state.line++
		// Ignore whitespace

	// string literal
	case '"':
		scanStringLiteral()
	default:
		if isDigit(char) {
			scanNumberLiteral()
		} else {
			error(scan_state.line, fmt.Sprintf("Unexpected character: %s", string(char)))
		}
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

// Return true and increment pointer if character at current pointer during scan is equal to the given rune
func match(next_char rune) bool {
	if scan_state.current >= len(source) {
		return false
	}
	if rune(source[scan_state.current]) != next_char {
		return false
	}

	scan_state.current++
	return true
}

// Peek the current pointer without moving the pointer
func peek() rune {
	if isAtEnd() { // fallback, should not reach here
		return rune(0)
	}
	return rune(source[scan_state.current])
}

func peekNext() rune {
	if scan_state.current+1 >= len(source) {
		return rune(0)
	}
	return rune(source[scan_state.current+1])
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func scanStringLiteral() {
	for {
		if isAtEnd() {
			error(scan_state.line, "Unterminated string.")
			return
		} else if peek() == '"' {
			advance()
			// without quotes
			addTokenWithLiteral(STRING, source[scan_state.start+1:scan_state.current-1])
			return
		} else if peek() == '\n' {
			scan_state.line++
		}

		advance()
	}
}

func scanNumberLiteral() {
	for isDigit(peek()) {
		advance()
	}

	if isDigit(peekNext()) && match('.') {
		for isDigit(peek()) {
			advance()
		}
	}

	num, err := strconv.ParseFloat(source[scan_state.start:scan_state.current], 64)
	if err != nil {
		error(scan_state.line, err.Error())
	}

	addTokenWithLiteral(NUMBER, num)
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
