package main

import (
	"fmt"
	"os"
)

// globals
var had_error bool = false

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
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
		run(string(file_contents))
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}

func run(source string) {
	if had_error {
		os.Exit(65)
	}

	tokens := scanTokens(source)

	for _, token := range tokens {
		fmt.Println(stringifyToken(token))
	}

}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error%s: %s", line, where, message)
}
