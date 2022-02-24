package main

import (
	"eslang/core"
	"eslang/interpreter"
)

func main() {
	var program core.Program

	program.Push(core.NewOP(core.OP_PUSH, 34))
	program.Push(core.NewOP(core.OP_PUSH, 35))
	program.Push(core.NewOP(core.OP_PLUS, nil))

	program.Push(core.NewOP(core.OP_DUMP, nil))

	program.Push(core.NewOP(core.OP_PUSH, 500))
	program.Push(core.NewOP(core.OP_PUSH, 80))
	program.Push(core.NewOP(core.OP_MINUS, nil))

	program.Push(core.NewOP(core.OP_DUMP, nil))

	interpreter.SimulateProgram(&program)
}
