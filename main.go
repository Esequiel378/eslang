package main

import (
	"eslang/core"
	"eslang/interpreter"
	"log"
)

func main() {
	program, err := core.NewProgramFromFile("./01-test-read-from-file.esl")

	if err != nil {
		log.Fatal(err)
	}

	interpreter.SimulateProgram(program)
}
