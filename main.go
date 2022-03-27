package main

import (
	"eslang/core"
	"eslang/interpreter"
	"flag"
	"log"
)

func main() {
	inputFile := flag.String("f", "./01-input.esl", ".esl file to run")
	flag.Parse()

	program := core.NewProgram()

	if err := core.LoadProgramFromFile(program, *inputFile); err != nil {
		log.Fatal(err)
	}

	if err := interpreter.SimulateProgram(program); err != nil {
		log.Fatal(err)
	}
}
