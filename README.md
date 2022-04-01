# Eslang

> Is a [Concatenative](https://en.wikipedia.org/wiki/Concatenative_programming_language) [Stack-Oriented](https://en.wikipedia.org/wiki/Stack-oriented_programming) Programming Language

## CLI

By default the CLI will read from a file called `01-input.esl`, but you can use the `-f` flag to specify another path

```bash
$ go run main.go -f test.esl
```

## Toc <a name="toc" />

+ [Push operations](#push-operations)
+ [Arithmetic operations](#arithmetic-operations)
+ [Stack operations](#stack-operations)

### Push operations <a name="push-operations" />

<br />

| **Operation**  | **Syntax**    | **Description**                            |
|----------------|---------------|--------------------------------------------|
| OP_PUSH_INT    | 378           | Pushes an 64 bytes int onto the stack      |
| OP_PUSH_FLOAT  | 3.78          | Pushes a 64 bytes float onto the stack     |
| OP_PUSH_BOOL   | true          | Pushes a boolean onto the stack            |
| OP_PUSH_STRING | "Hello World" | Pushes a string onto the stack             |

> Large integer numbers can divided by an undercore `_` to make them more readable: `1_000`, `100_000`

<small>[Top ▲](#toc)</small>

### Arithmetic operations <a name="arithmetic-operations" />

<br />

| **Operation**   | **Syntax** | **Description**                                                               |
|-----------------|:----------:|-------------------------------------------------------------------------------|
| OP_OPERATOR_ADD | +          | Push the addition of the two topmost values on the stack onto the stack       |
| OP_OPERATOR_SUB | -          | Push the subtraction of the two topmost values on the stack onto the stack    |
| OP_OPERATOR_MUL | *          | Push the multiplication of the two topmost values on the stack onto the stack |
| OP_OPERATOR_DIV | /          | Push the division of the two topmost values on the stack onto the stack       |
| OP_OPERATOR_MOD | %          | Push the modulo of the two topmost values on the stack onto the stack         |

<small>[Top ▲](#toc)</small>

### Stack operations <a name="stack-operations" />

<br />

| **Operation** | **Syntax** | **Examples**                 | **Description**                                                       |
|---------------|------------|------------------------------|-----------------------------------------------------------------------|
| TOKEN_DROP    | drop       | ( a -- )                     | Drops the top of the stack                                            |
| TOKEN_DUMP    | dump       | ( a -- )                     | Dumps the stack                                                       |
| TOKEN_DUP     | dup        | ( a -- a a )                 | Duplicates the top of the stack                                       |
| TOKEN_OVER    | over       | ( a b -- a b a )             | Duplicate the second-to-top of the stack                              |
| TOKEN_ROT     | rot        | ( a b c -- b c a )           | Rotates the top three elements of the stack                           |
| TOKEN_O_ROT   | -rot       | ( a b c -- c a b ) rot rot   | Rotates the top three elements of the stack in the opposite direction |
| TOKEN_SWAP    | swap       | ( a b -- b a )               | Swaps the top two elements of the stack                               |
| TOKEN_TUCK    | tuck       | ( a b -- b a b ) swap over   | Duplicates the top of the stack and places it below the second-to-top |

<small>[Top ▲](#toc)</small>
