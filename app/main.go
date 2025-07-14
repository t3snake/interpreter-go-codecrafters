package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/app/loxerrors"
	"github.com/codecrafters-io/interpreter-starter-go/app/parser"
	"github.com/codecrafters-io/interpreter-starter-go/app/scanner"
	"github.com/codecrafters-io/interpreter-starter-go/app/token"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	valid_commands := []string{"tokenize", "parse"}

	is_command_valid := false
	for _, valid_command := range valid_commands {
		if command == valid_command {
			is_command_valid = true
		}
	}

	if !is_command_valid {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	file_contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(file_contents) > 0 {
		run(command, string(file_contents))
	} else {
		fmt.Println("EOF  null")
	}
}

func run(command, source string) {
	had_error := loxerrors.GetErrorState()

	tokens := scanner.ScanTokens(source)
	if command == "tokenize" {
		for _, token_ := range tokens {
			fmt.Println(token.StringifyToken(token_))
		}

		if *had_error {
			os.Exit(65)
		}

		return // early return if only to tokenize
	}

	ast, _ := parser.Parse(tokens)

	if *had_error {
		os.Exit(65)
	}

	if command == "parse" {
		fmt.Println(parser.AstPrinter(ast))
	}
}
