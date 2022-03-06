package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TwiN/go-color"
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
		if len(strings.Trim(line, " ")) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")

		for _, token := range tokens {
			token = strings.Trim(token, " ")

			if len(token) == 0 {
				continue
			}

			found := false

			for _, tokenHandler := range REGISTERED_TOKENS {
				operation, err := tokenHandler(token, line, lnum, &blocks)

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
				cnum := strings.Index(line, token)

				return fmt.Errorf(
					"Token error in %d:%d - '%s' is not a valid token.",
					lnum+1, cnum+1,
					token,
				)
			}

		}
	}

	if !blocks.IsEmpty() {
		block := blocks.Last()

		tokenStart := block.TokenStart().TokenAlias()
		tokenEnd := block.TokenEnd().TokenAlias()
		lnum, cnum := block.TokenStart().Position()

		return fmt.Errorf(
			"%s missing %s closing token in line %d:%d",
			color.InYellow(tokenStart),
			color.InYellow(tokenEnd),
			lnum, cnum,
		)
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
