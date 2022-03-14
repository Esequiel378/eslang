package core

import "fmt"

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("can not perform `Stack.Pop()`, stack is empty")
	}

	// Get the index of the top most element.
	index := len(*s) - 1
	// Index into the slice and obtain the element.
	value := (*s)[index]
	// Remove it from the stack by slicing it off.
	*s = (*s)[:index]

	return value, nil
}

func (s *Stack) PopTwo() (lhs interface{}, rhs interface{}, err error) {
	rhs, err = s.Pop()

	if err != nil {
		return nil, nil, err
	}

	lhs, err = s.Pop()

	if err != nil {
		return nil, nil, err
	}

	return lhs, rhs, nil
}

type BlockStack []BlockOperation

func (bs *BlockStack) IsEmpty() bool {
	return len(*bs) == 0
}

func (bs *BlockStack) Push(value BlockOperation) {
	*bs = append(*bs, value)
}

func (bs *BlockStack) Tail() BlockOperation {
	return (*bs)[len(*bs)-1]
}

func (bs *BlockStack) Pop() (BlockOperation, error) {
	if bs.IsEmpty() {
		return nil, fmt.Errorf("can not perform `ProgramStack.Pop()`, stack is empty")
	}

	// Get the index of the top most element.
	index := len(*bs) - 1
	// Index into the slice and obtain the element.
	value := (*bs)[index]
	// Remove it from the stack by slicing it off.
	*bs = (*bs)[:index]

	return value, nil
}
