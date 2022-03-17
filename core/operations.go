package core

type OperationType int

const (
	OP_BLOCK OperationType = iota
	OP_DUMP                = iota
	OP_MOP                 = iota
	OP_PUSH                = iota
)

type Block struct {
	current    *Program
	next       *Block
	tokenStart *Token
	tokenEnd   *Token
}

func NewEmptyBlock() *Block {
	return &Block{
		current:    &Program{},
		next:       &Block{},
		tokenStart: &Token{},
		tokenEnd:   &Token{},
	}
}

func NewBlock(tokenStart, tokenEnd TokenType) *Block {
	return &Block{
		current:    &Program{},
		next:       &Block{},
		tokenStart: &Token{token: tokenStart},
		tokenEnd:   &Token{token: tokenEnd},
	}
}

func (b *Block) Current() *Program {
	return b.current
}

func (b *Block) SetNext(next *Block) {
	b.next = next
}

func (b *Block) HasNext() bool {
	return b.next != nil
}

func (b *Block) Next() *Block {
	return b.next
}

func (b *Block) Last() *Block {
	if b.HasNext() {
		return b.Next().Last()
	}

	return b
}

func (b *Block) IsClose() bool {
	return !b.IsOpen()
}

func (b *Block) IsOpen() bool {
	line, col := b.tokenEnd.Position()

	return line == 0 && col == 0
}

func (b *Block) TokenStart() *Token {
	return b.tokenStart
}

func (b *Block) TokenEnd() *Token {
	return b.tokenEnd
}

type Type int

const (
	Nil   Type = iota
	Int        = iota
	Float      = iota
)

type OperationValue struct {
	intValue   int64
	floatValue float64
	block      *Block
	_type      Type
}

func NewOperationValue() *OperationValue {
	return &OperationValue{
		intValue:   0,
		floatValue: 0,
		block:      NewEmptyBlock(),
		_type:      Nil,
	}
}

func (o *OperationValue) Type() Type {
	return o._type
}

func (o *OperationValue) Int() int64 {
	return o.intValue
}

func (o *OperationValue) SetInt(value int64) *OperationValue {
	o.intValue = value
	o._type = Int

	return o
}

func (o *OperationValue) Float() float64 {
	return o.floatValue
}

func (o *OperationValue) SetFloat(value float64) *OperationValue {
	o.floatValue = value
	o._type = Float

	return o
}

func (o *OperationValue) Block() *Block {
	return o.block
}

type Operation struct {
	opType     OperationType
	opValue    *OperationValue
	tokenStart *Token
	tokenEnd   *Token
}

func NewOperation(op OperationType, value *OperationValue, tokenStart, tokenEnd TokenType) *Operation {
	return &Operation{
		opType:     op,
		opValue:    value,
		tokenStart: &Token{token: tokenStart},
		tokenEnd:   &Token{token: tokenEnd},
	}
}

func (o Operation) Type() OperationType {
	return o.opType
}

func (o Operation) Value() *OperationValue {
	return o.opValue
}

func (o Operation) TokenStart() *Token {
	return o.tokenStart
}

func (o Operation) TokenEnd() *Token {
	return o.tokenEnd
}
