package core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Program []*Operation

func (p *Program) IsEmpty() bool {
	return len(*p) == 0
}

func (p *Program) Push(operation *Operation) {
	*p = append(*p, operation)
}

func (p *Program) Pop() (*Operation, error) {
	if p.IsEmpty() {
		return nil, fmt.Errorf("can not perform `Program.Pop()`, program is empty.")
	}

	// Get the index of the top most element.
	index := len(*p) - 1
	// Index into the slice and obtain the element.
	operation := (*p)[index]
	// Remove it from the stack by slicing it off.
	*p = (*p)[:index]

	return operation, nil
}

func NewProgramFromFile(filename string) (*Program, error) {
	lines, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var program Program

	for lineNumber, line := range strings.Split(string(lines), "\n") {
		for _, token := range strings.Split(line, " ") {
			found := false

			for _, tokenHandler := range REGISTERED_TOKENS {
				operation, err := tokenHandler(token)

				if err != nil {
					continue
				}

				program.Push(operation)
				found = true
				break
			}

			if !found {
				log.Fatal(
					fmt.Errorf("Token error ln:%d/%d - %s is not recognized as a valid token.",
						lineNumber, strings.Index(line, token)+1, token,
					),
				)
			}
		}
	}

	return &program, nil
}
