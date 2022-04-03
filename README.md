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
+ [Relational operations](#relational-operations)
+ [Logical operations](#logical-operations)
+ [Stack operations](#stack-operations)
+ [Blocks](#blocks)
    + [Conditions](#conditions)
    + [Loops](#loops)
+ [Variables](#variables)

### Push operations <a name="push-operations" />

<br />

| **Operation**  | **Syntax**    | **Description**                            |
|----------------|---------------|--------------------------------------------|
| OP_PUSH_INT    | 378           | Push a 64 bytes int onto the stack      |
| OP_PUSH_FLOAT  | 3.78          | Push a 64 bytes float onto the stack     |
| OP_PUSH_BOOL   | true          | Push a boolean onto the stack            |
| OP_PUSH_STRING | "Hello World" | Push a string onto the stack             |

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

### Relational operations <a name="relational-operations" />

<br />


| Operation                           | Syntax | Description                                                                                           |
|-------------------------------------|:------:|-------------------------------------------------------------------------------------------------------|
| OP_R_OPERATOR_EQUAL                 | =      | Push the equal comparition between the two topmost values on the stack onto the stack                 |
| OP_R_OPERATOR_NOT_EQUAL             | !=     | Push the not equal comparition between the two topmost values on the stack onto the stack             |
| OP_R_OPERATOR_LESS_THAN             | <      | Push the less than comparition between the two topmost values on the stack onto the stack             |
| OP_R_OPERATOR_LESS_THAN_OR_EQUAL    | <=     | Push the less than or equal comparition between the two topmost values on the stack onto the stack    |
| OP_R_OPERATOR_GREATER_THAN          | >      | Push the greater than comparition between the two topmost values on the stack onto the stack          |
| OP_R_OPERATOR_GREATER_THAN_OR_EQUAL | >=     | Push the greater than or equal comparition between the two topmost values on the stack onto the stack |

<small>[Top ▲](#toc)</small>

### Logical operations <a name="logical-operations" />

<br />

| Operation         | Syntax | Description                                                                     |
|-------------------|:------:|---------------------------------------------------------------------------------|
| OP_L_OPERATOR_AND | &&     | Push the logical or between the two topmost values on the stack onto the stack  |
| OP_L_OPERATOR_NOT | !      | Push the logical and between the two topmost values on the stack onto the stack |
| OP_L_OPERATOR_OR  | \|\|   | Push the logical not between the two topmost values on the stack onto the stack |

<small>[Top ▲](#toc)</small>

### Stack operations <a name="stack-operations" />

<br />

| **Operation** | **Syntax** | **Examples**                 | **Description**                                                       |
|---------------|------------|------------------------------|-----------------------------------------------------------------------|
| OP_DROP       | drop       | ( a -- )                     | Drops the top of the stack                                            |
| OP_DUMP       | dump       | ( a -- )                     | Dumps the stack                                                       |
| OP_DUP        | dup        | ( a -- a a )                 | Duplicates the top of the stack                                       |
| OP_NIP        | nip        | ( a b -- b ) swap drop       | Drops the second-to-top element of the stack                          |
| OP_OVER       | over       | ( a b -- a b a )             | Duplicate the second-to-top of the stack                              |
| OP_O_ROT      | rot        | ( a b c -- b c a )           | Rotates the top three elements of the stack                           |
| OP_ROT        | -rot       | ( a b c -- c a b ) rot rot   | Rotates the top three elements of the stack in the opposite direction |
| OP_SWAP       | swap       | ( a b -- b a )               | Swaps the top two elements of the stack                               |
| OP_TUCK       | tuck       | ( a b -- b a b ) swap over   | Duplicates the top of the stack and places it below the second-to-top |
| OP_TWO_DROP   | 2drop      | ( a b -- )                   | Drops the top two items from the stack 
| OP_TWO_DUP    | 2dup       | ( a b -- a b a b ) over over | Duplicates the top two elements of the stack                          |
| OP_TWO_OVER   | 2over      | ( a b c d -- a b c d a b )   | Duplicates the second-to-top two items to the top of the stack 
| OP_TWO_SWAP   | 2swap      | ( a b c d -- c d a b )       | Swaps the second-to-top two items on the stack 

<small>[Top ▲](#toc)</small>

### Blocks <a name="blocks" />

#### Conditions <a name="conditions" />

> For now, Eslang have `if-else` conditions, In the future, I'm plaing to add `switch` conditions,
but I have no itention to add `if-elif-else` conditions. This decision is maily to keep the language
as simple as posible.

#### If-else

```pascal
true if
    "Hello world from if condition" dump
end
```

<small>[Top ▲](#toc)</small>

#### Loops <a name="loops" />

##### While

```pascal
10 while dup > 0 do
    dup dump
    1 -
end
```

<small>[Top ▲](#toc)</small>

### Variables <a name="variables" />

##### To use a variable, we need to give it a name and a type, with following syntax

```pascal
counter int
```

> Where `counter` is the variable name and `int` is the variable type

##### By default variables will have a zero value.

| Type  | Zero value |
|-------|------------|
| int   | 0          |
| float | 0.0        |
| bool  | false      |
| str   | ""         |

##### Variables operations

| Operation         | Syntax  | Example      | Description                                                                   |
|-------------------|---------|--------------|-------------------------------------------------------------------------------|
| OP_VARIABLE       | counter | counter dump | Create a variable if it doesn't exist and push it to the stack                |
| OP_VARIABLE_WRITE | .       | counter 1 .  | Write the value top of the stack to a variable in the second top of the stack |
