package main

import (
	"fmt"
	"log"
	"strings"
)

import (
	cli "github.com/Hilson-Alex/go-tour-cli/command"
	"github.com/Hilson-Alex/go-tour-cli/outputFormat"
)

var exampleMap = cli.ExampleMap

func main() {
	var value string
	log.SetPrefix("\033[91m")
	for keep := true; keep; keep = evaluateCommand(value) {
		cli.PrintHeader("EXAMPLE")
		fmt.Printf("Type \"%s\" to see the options or \"quit\" to quit the program.\n", cli.HELP)
		fmt.Print("example -> ")
		if _, err := fmt.Scanln(&value); err != nil {
			log.Fatal(err)
		}
	}
}

func evaluateCommand(command string) (keep bool) {
	fmt.Println()
	example, keep := exampleMap[cli.Prompt(strings.ToLower(command))]
	if !keep {
		if command != "quit" {
			fmt.Println(outputFormat.AsError("Exercise not found. Quiting program"))
		}
		return
	}
	example()
	return
}
