package core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Program []Operation

func (p *Program) IsEmpty() bool {
	return len(*p) == 0
}

func (p *Program) Push(operation Operation) {
	*p = append(*p, operation)
}

func (p *Program) parseLines(lines []string) error {
	var blocks BlockStack

	for lnum, line := range lines {
		line = strings.Trim(line, " ")

		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")

		for cnum, token := range tokens {
			token = strings.Trim(token, " ")
			found := false

			for _, tokenHandler := range REGISTERED_TOKENS {
				operation, err := tokenHandler(token, &blocks)

				if err != nil {
					continue
				}

				found = true

				if operation == nil {
					break
				}

				if blocks.IsEmpty() {
					p.Push(operation)
				} else {
					blocks.Last().PushIntoBlocks(operation)
				}

				break
			}

			if !found {
				return fmt.Errorf(
					"Token error in %d:%d - '%s' is not a valid token.",
					lnum+1, cnum+1,
					token,
				)
			}

		}
	}

	return nil
}

func (p *Program) LoadFromFile(filename string) error {
	rawLines, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	lines := strings.Split(string(rawLines), "\n")

	if err := p.parseLines(lines); err != nil {
		log.Fatal(err)
	}

	// fmt.Println((*program)[1])

	return nil
}
