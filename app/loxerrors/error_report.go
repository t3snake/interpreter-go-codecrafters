package loxerrors

import (
	"fmt"
	"os"

	. "github.com/codecrafters-io/interpreter-starter-go/app/token"
)

var Had_error bool = false

func Scanner_error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	Had_error = true
}

func Parser_error(token Token, message string) {
	if token.Type == EOF {
		report(token.Line, " at end", message)
	} else {
		report(token.Line, " at '"+token.Lexeme+"'", message)
	}
}
