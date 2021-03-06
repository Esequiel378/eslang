package main

import (
	ops "eslang/core/operations"
	"eslang/core/utils"
	"eslang/interpreter"
	"flag"
	"log"
)

func main() {
	inputFile := flag.String("f", "./01-input.esl", ".esl file to run")
	flag.Parse()

	program := ops.NewProgram(*inputFile)

	if err := utils.LoadProgramFromFile(program, *inputFile); err != nil {
		log.Fatal(err)
	}

	if err := interpreter.RunProgram(program); err != nil {
		log.Fatal(err)
	}
}
