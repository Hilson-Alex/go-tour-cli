// Author: [Hilson-Alex](https://github.com/Hilson-Alex)

package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt is a type for user interaction on the CLI.
type Prompt string

// Enum for the prompts of each example.
const (
	SQRT       Prompt = "sqrt"
	IMAGE      Prompt = "image"
	WORD_COUNT Prompt = "word_count"
	FIBONACCI  Prompt = "fib"
	ROT13      Prompt = "rot13"
	HELP       Prompt = "help"
)

// describe is a method to give an explanation for a function.
func (prompt Prompt) describe(description string) {
	prompt.printCursor()
	fmt.Println(description)
}

// printCursor prints a signal of the prompt, usually for data input.
func (prompt Prompt) printCursor() {
	fmt.Printf("%s -> ", prompt)
}

// ReadValue takes a user input and load it on the response, casting
// its type.
// Breaks space separated values in multiple responses, so you should
// use ReadString if you need to keep the spaces.
func ReadValue[T any](prompt Prompt, response *T) error {
	prompt.printCursor()
	var _, err = fmt.Scanln(response)
	return err
}

// ReadString reads a user input as text.
func (prompt Prompt) ReadString() string {
	var scanner = bufio.NewScanner(os.Stdin)
	prompt.printCursor()
	scanner.Scan()
	return scanner.Text()
}

// AskRerun asks the user for keeping on the exercise
func (prompt Prompt) AskRerun() bool {
	fmt.Println()
	defer fmt.Println()
	fmt.Println("Do you want to run the exercise again? Y/N")
	var response = prompt.ReadString()
	var firstChar = []rune(strings.ToLower(response))[0]
	if firstChar == 'y' {
		return true
	}
	return false
}
