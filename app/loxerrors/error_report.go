package loxerrors

import (
	"fmt"
	"os"
	"sync"

	//lint:ignore ST1001 I dont care
	. "github.com/codecrafters-io/interpreter-starter-go/app/token"
)

type Singleton struct {
	had_error bool
}

var instance *Singleton
var once sync.Once

func GetErrorState() *bool {
	once.Do(func() {
		instance = &Singleton{had_error: false}
	})

	return &instance.had_error
}

func Scanner_error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	instance.had_error = true
}

func Parser_error(token Token, message string) {
	if token.Type == EOF {
		report(token.Line, " at end", message)
	} else {
		report(token.Line, " at '"+token.Lexeme+"'", message)
	}
}
