Interpreter for the [Monkey programming language](https://monkeylang.org/), written in Go.

## Features

The Monkey language interpreter supports:

- Integer and boolean data types
- String data types
- Array and hash data types
- Arithmetic expressions
- Logical expressions
- If-else expressions
- Functions and higher-order functions
- Closures
- Built-in functions
- Macro system

## Examples

Here are some examples of what you can do with the Monkey language:

### Arithmetic Expressions

```monkey
let a = 5;
let b = 10;
let sum = a + b; // 15
```

### Logical Expressions

```monkey
let a = true;
let b = false;
let result = a && b; // false
```

### If-Else Expressions

```monkey
let a = 10;
if (a > 5) {
    puts("a is greater than 5");
} else {
    puts("a is not greater than 5");
}
```

### Functions and Higher-Order Functions

```monkey
let add = fn(a, b) { return a + b; };
let result = add(5, 10); // 15

let applyFunc = fn(a, b, func) { return func(a, b); };
let result = applyFunc(5, 10, add); // 15
```

### Closures

```monkey
let newAdder = fn(a) { return fn(b) { return a + b; }; };
let addTwo = newAdder(2);
let result = addTwo(2); // 4
```

### Built-In Functions

```monkey
let arr = [1, 2, 3, 4, 5];
let len = len(arr); // 5
```

### Macro System

```monkey
let unless = macro(condition, consequence, alternative) {
    quote(if (!(unquote(condition))) {
        unquote(consequence);
    } else {
        unquote(alternative);
    });
};

unless(10 > 5, puts("not greater"), puts("greater"));
```

## Running the Interpreter

You can run the interpreter in two ways:

1. Interactive mode: Run the interpreter without any arguments to start a REPL (Read-Eval-Print Loop), where you can enter Monkey language expressions interactively.

```bash
go run main.go
```

2. File mode: Run the interpreter with a file name as an argument to interpret the Monkey language code in the file.

```bash
go run main.go examples/example.monkey
```