## Program initialization and execution ## {#Program_initialization_and_execution}

### The zero value ### {#The_zero_value}

When storage is allocated for a [variable](#Variables), either through a declaration or a call of `new`, or when a new value is created, either through a composite literal or a call of `make`, and no explicit initialization is provided, the variable or value is given a default value. Each element of such a variable or value is set to the _zero value_ for its type: `false` for booleans, `0` for integers, `0.0` for floats, `""` for strings, and `nil` for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.

These two simple declarations are equivalent:

``` go
var i int
var i int = 0
```

After

``` go
type T struct { i int; f float64; next *T }
t := new(T)
```

the following holds:

``` go
t.i == 0
t.f == 0.0
t.next == nil
```

The same would also be true after

``` go
var t T
```

### Package initialization ### {#Package_initialization}

Within a package, package-level variables are initialized in _declaration order_ but after any of the variables they _depend_ on.

More precisely, a package-level variable is considered _ready for initialization_ if it is not yet initialized and either has no [initialization expression](#Variable_declarations) or its initialization expression has no dependencies on uninitialized variables. Initialization proceeds by repeatedly initializing the next package-level variable that is earliest in declaration order and ready for initialization, until there are no variables ready for initialization.

If any variables are still uninitialized when this process ends, those variables are part of one or more initialization cycles, and the program is not valid.

The declaration order of variables declared in multiple files is determined by the order in which the files are presented to the compiler: Variables declared in the first file are declared before any of the variables declared in the second file, and so on.

Dependency analysis does not rely on the actual values of the variables, only on lexical _references_ to them in the source, analyzed transitively. For instance, if a variable `x`'s initialization expression refers to a function whose body refers to variable `y` then `x` depends on `y`. Specifically:

*   A reference to a variable or function is an identifier denoting that variable or function.
*   A reference to a method `m` is a [method value](#Method_values) or [method expression](#Method_expressions) of the form `t.m`, where the (static) type of `t` is not an interface type, and the method `m` is in the [method set](#Method_sets) of `t`. It is immaterial whether the resulting function value `t.m` is invoked.
*   A variable, function, or method `x` depends on a variable `y` if `x`'s initialization expression or body (for functions and methods) contains a reference to `y` or to a function or method that depends on `y`.

Dependency analysis is performed per package; only references referring to variables, functions, and methods declared in the current package are considered.

For example, given the declarations

``` go
var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}
```

the initialization order is `d`, `b`, `c`, `a`.

Variables may also be initialized using functions named `init` declared in the package block, with no arguments and no result parameters.

``` go
func init() { … }
```

Multiple such functions may be defined per package, even within a single source file. In the package block, the `init` identifier can be used only to declare `init` functions, yet the identifier itself is not [declared](#Declarations_and_scope). Thus `init` functions cannot be referred to from anywhere in a program.

A package with no imports is initialized by assigning initial values to all its package-level variables followed by calling all `init` functions in the order they appear in the source, possibly in multiple files, as presented to the compiler. If a package has imports, the imported packages are initialized before initializing the package itself. If multiple packages import a package, the imported package will be initialized only once. The importing of packages, by construction, guarantees that there can be no cyclic initialization dependencies.

Package initialization—variable initialization and the invocation of `init` functions—happens in a single goroutine, sequentially, one package at a time. An `init` function may launch other goroutines, which can run concurrently with the initialization code. However, initialization always sequences the `init` functions: it will not invoke the next one until the previous one has returned.

To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.

### Program execution ### {#Program_execution}

A complete program is created by linking a single, unimported package called the _main package_ with all the packages it imports, transitively. The main package must have package name `main` and declare a function `main` that takes no arguments and returns no value.

``` go
func main() { … }
```

Program execution begins by initializing the main package and then invoking the function `main`. When that function invocation returns, the program exits. It does not wait for other (non-`main`) goroutines to complete.
