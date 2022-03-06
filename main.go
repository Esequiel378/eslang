package main

import (
	"eslang/core"
	"eslang/interpreter"
	"flag"
	"log"
)

func main() {
	inputFile := flag.String("f", "./01-input.esl", ".esl file to run")
	mode := flag.String("m", "i", "Mode in which theprogram will be executed")
	flag.Parse()

	var program core.Program

	if err := program.LoadFromFile(*inputFile); err != nil {
		log.Fatal(err)
	}

	switch *mode {
	case "i":
		if err := interpreter.SimulateProgram(&program); err != nil {
			log.Fatal(err)
		}
	case "v":
		if err := core.PrintProgram(&program, 0); err != nil {
			log.Fatal(err)
		}
	}

}
