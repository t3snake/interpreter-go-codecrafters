package main

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

	default:
		error(scan_state.line, "Unexpected character.")
	}
}

func isAtEnd() bool {
	return scan_state.current >= len(source)
}

func advance() rune {
	scan_state.current++
	return rune(source[scan_state.current-1])
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
