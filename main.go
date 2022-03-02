package main

import (
	"eslang/core"
	"eslang/interpreter"
	"log"
)

func main() {
	var program core.Program

	if err := program.LoadFromFile("./01-test-read-from-file.esl"); err != nil {
		log.Fatal(err)
	}

	if err := interpreter.SimulateProgram(&program); err != nil {
		log.Fatal(err)
	}
}
