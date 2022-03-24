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

### ToC

+ [Store values on the stack](#push)
+ [Strings literals](#string-literals)
+ [Dump](#dump)
+ [Dup](#dup)
+ [Operations](#operations)
+ [Variables](#variables)
    - [write](#write)
    - [read](#read)
+ [Blocks](#blocks)
    - [`DO-END`](#do-end)
    - [`IF-ELSE-END`](#if-else-end)
+ [While](#While)

### Store values in the stack <a name="push" />

Input

```pascal
1 2 3 4
```

<details>
    <summary>Stack tree</summary>
  
```pascal
PUSH_INT 1 in line 1:1
PUSH_INT 2 in line 1:3
PUSH_INT 3 in line 1:5
PUSH_INT 4 in line 1:7
```
</details>

### String Literals <a name="string-literals" />

Input

```pascal
"Hello world"
```

<details>
  <summary>Stack tree</summary>
  
```pascal
PUSH_STR "Hello world" in line 0:1
```
</details>

### Dump last value from the stack <a name="dump" />

Input

```pascal
1 2 dump
3 4 dump dump
```

Output

```pascal
2
4
3
```

<details>
  <summary>Stack tree</summary>
  
```pascal
PUSH_INT 1 in line 1:1
PUSH_INT 2 in line 1:3
DUMP in line 1:5
PUSH_INT 3 in line 2:1
PUSH_INT 4 in line 2:3
DUMP in line 2:5
DUMP in line 2:5
```
</details>

### Dup last value from the stack <a name="dup" />

Input

```pascal
2 dup dump dump
```

Output

```pascal
2
2
```

<details>
  <summary>Stack tree</summary>
  
```pascal
PUSH_INT 2 in line 1:1
DUP in line 1:3
DUMP in line 1:7
DUMP in line 1:7
```
</details>

### Operations <a name="operations" />

Input

```pascal
1 2 + dump
3 2 - dump
"Hello " "World" + dump
```

Output

```pascal
3
1
Hello World
```

<details>
  <summary>Stack tree</summary>
  
```pascal
PUSH_INT 1 in line 1:1
PUSH_INT 2 in line 1:3
PLUS in line 1:5
DUMP in line 1:7
PUSH_INT 3 in line 2:1
PUSH_INT 2 in line 2:3
MINUS in line 2:5
DUMP in line 2:7
PUSH_STR Hello  in line 3:1
PUSH_STR World in line 3:10
PLUS in line 3:18
DUMP in line 3:20
```
</details>

### Variables <a name="variables" />

> Declare variable using a name and a type. Name must be in camel case

Input

```pascal
tmp int 
```

<details>
  <summary>Stack tree</summary>
  
```pascal
VAR in line 1:1
```
</details>


#### Write <a name="write" />

`.` (dot) takes a value, a variable, and write it into the stack memory. Pushes the stored value into the top of the stack

Input

```pascal
tmp int
1 tmp .
```

<details>
  <summary>Stack tree</summary>
  
```pascal
VAR in line 1:1
PUSH_INT 1 in line 2:1
VAR in line 2:3
VAR_WRITE in line 2:7
```
</details>

> Keep in mind that operations with variables will return the result of such operation, and no the pointer to the variable.

Won't work

```pascal
11 tmp int .
tmp 1 - .
dump
```

Will work

```pascal
11 tmp int .
tmp 1 - tmp .
dump
```

Output

```
10
```

#### Read <a name="read" />

`,` (comma) take a variable name and push the value into the top of the stack.
> There is a shorthand for reading a variable, by using the name without the comma `tmp dump`.

Input

```pascal
1 tmp int .
tmp , dump 
// tmp dump
```

Output

```pascal
1
```


<details>
  <summary>Stack tree</summary>
  
```pascal
PUSH_INT 1 in line 1:1
VAR in line 1:3
VAR_WRITE in line 1:11
VAR in line 2:1
VAR_READ in line 2:5
DUMP in line 2:7
```
</details>


### Blocks <a name="blocks" />


#### do-end <a name="do-end" />

Input

```pascal
do
    1 2 + dump
end
```

Output

```
3
```

<details>
  <summary>Stack tree</summary>
  
```pascal
DO in lines [1:1:3:1]
        PUSH_INT 1 in line 2:1
        PUSH_INT 2 in line 2:3
        PLUS in line 2:5
        DUMP in line 2:7
END in line 3:1
```
</details>

#### if-else-end <a name="if-else-end" />

Input

```pascal
1 if
    1 dump
else
    0 dump
end
```

Output

```pascal
1
```

<details>
  <summary>Stack tree</summary>
  
```pascal
PUSH_INT 1 in line 1:1
IF in lines [1:3:5:1]
        PUSH_INT 1 in line 2:1
        DUMP in line 2:3
ELSE in lines [1:3:5:1]
        PUSH_INT 0 in line 4:1
        DUMP in line 4:3
END in line 5:1
```
</details>

#### While <a name="while" />

Input

```pascal
5 while
    dup dump
    1 -
end
```

Output

```pascal
5
4
3
2
1
```

<details>
  <summary>Stack tree</summary>
  
```pascal
PUSH_INT 5 in line 1:1
WHILE in lines [1:3:4:1]
        DUP in line 2:5
        DUMP in line 2:9
        PUSH_INT 1 in line 3:5
        MINUS in line 3:7
END in line 4:1
```
</details>
