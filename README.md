# Eslang

> Is a [Concatenative](https://en.wikipedia.org/wiki/Concatenative_programming_language) [Stack-Oriented](https://en.wikipedia.org/wiki/Stack-oriented_programming) Programming Language

### CLI

eslang has two modes to run

1. Interpret mode `-i`

> Will interpret the program and output the results

```bash
$ go run main.go -m i
```

2. Visual mode `-v`

> Will print the program stack tree

```bash
$ go run main.go -m v
```

By default the CLI will read from a file called `01-input.esl`, but you can use the `-f` flag to specify another path

```bash
$ go run main.go -m v -r test.esl
```

## Examples

### TOC

+ [Store values on the stack](#push)
+ [Dump](#dump)
+ [Operations](#operations)
+ [Blocks](#blocks)
    - [`DO-END`](#do-end)
    - [`IF-ELSE-END`](#if-else-end)

### Store values in the stack (only numbers for now) <a name="push" />

Input

```pascal
1 2 3 4
```

Output

```pascal
```

Stack tree

```pascal
PUSH 1 in line 1:1
PUSH 2 in line 1:3
PUSH 3 in line 1:5
PUSH 4 in line 1:7
```

### Dump last value from the stack <a name="dump" />

Input

```pascal
1 2 .
3 4 . .
```

Output (The new line character is a bug, will be fixed soon)

```pascal
2
3
4
```

Intended Output

```pascal
234
```

Stack tree

```pascal
PUSH 1 in line 1:1
PUSH 2 in line 1:3
DUMP in line 1:5
PUSH 3 in line 2:1
PUSH 4 in line 2:3
DUMP in line 2:5
DUMP in line 2:5
```

### Operations <a name="operations" />

Input

```pascal
1 2 + .
3 2 - .
```

Output

```pascal
3
1
```

Stack tree

```pascal
PUSH 1 in line 1:1
PUSH 2 in line 1:3
PLUS in line 1:5
DUMP in line 1:7
PUSH 3 in line 2:1
PUSH 2 in line 2:3
MINUS in line 2:5
DUMP in line 2:7
```

### Blocks <a name="block" />


#### `DO-END` <a name="do-end" />

Input

```pascal
do
    1 2 + .
end
```

Output

```
3
```

Stack tree

```pascal
DO in lines [1:1:3:1]
        PUSH 1 in line 2:5
        PUSH 2 in line 2:7
        PLUS in line 2:9
        DUMP in line 2:11
END in line 3:1
```

#### `IF-ELSE-END` <a name="if-else-end" />

Input

```pascal
1 if
    1 .
else
    0 .
end
```

Output

```pascal
1
```

Stack tree

```pascal
PUSH 1 in line 1:1
IF in lines [1:3:3:1]
        PUSH 1 in line 2:5
        DUMP in line 2:7
ELSE in lines [1:3:5:1]
        PUSH 0 in line 4:5
        DUMP in line 4:7
END in line 5:1
```
