package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"brainfuck/brainfuck"
	"io"
)

func main() {
	if len(os.Args) <= 1 {
		errorAndDie(os.Stderr, "%s\n", "Usage: brainfuck <filepath>")
	}
	code, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		errorAndDie(os.Stderr, "Error: %s\n", err)
	}

	c := brainfuck.NewCompiler(string(code))
	instructions := c.Compile()

	m := brainfuck.NewMachine(instructions, os.Stdin, os.Stdout)
	m.Execute()
}

func errorAndDie(w io.Writer, format string, lit interface{}) {
	fmt.Fprintf(w, format, lit)
	os.Exit(-1)
}