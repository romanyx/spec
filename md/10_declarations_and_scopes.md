## Declarations and scope ## {#Declarations_and_scope}

A _declaration_ binds a non-[blank](#Blank_identifier) identifier to a [constant](#Constant_declarations), [type](#Type_declarations), [variable](#Variable_declarations), [function](#Function_declarations), [label](#Labeled_statements), or [package](#Import_declarations). Every identifier in a program must be declared. No identifier may be declared twice in the same block, and no identifier may be declared in both the file and package block.

The [blank identifier](#Blank_identifier) may be used like any other identifier in a declaration, but it does not introduce a binding and thus is not declared. In the package block, the identifier `init` may only be used for [`init` function](#Package_initialization) declarations, and like the blank identifier it does not introduce a new binding.

<pre class="ebnf"><a id="Declaration">Declaration</a>   = <a href="#ConstDecl" class="noline">ConstDecl</a> | <a href="#TypeDecl" class="noline">TypeDecl</a> | <a href="#VarDecl" class="noline">VarDecl</a> .
<a id="TopLevelDecl">TopLevelDecl</a>  = <a href="#Declaration" class="noline">Declaration</a> | <a href="#FunctionDecl" class="noline">FunctionDecl</a> | <a href="#MethodDecl" class="noline">MethodDecl</a> .
</pre>

The _scope_ of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, label, or package.

Go is lexically scoped using [blocks](#Blocks):

1.  The scope of a [predeclared identifier](#Predeclared_identifiers) is the universe block.
2.  The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block.
3.  The scope of the package name of an imported package is the file block of the file containing the import declaration.
4.  The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.
5.  The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.
6.  The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

The [package clause](#Package_clause) is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same [package](#Packages) and to specify the default package name for import declarations.

### Label scopes ### {#Label_scopes}

Labels are declared by [labeled statements](#Labeled_statements) and are used in the ["break"](#Break_statements), ["continue"](#Continue_statements), and ["goto"](#Goto_statements) statements. It is illegal to define a label that is never used. In contrast to other identifiers, labels are not block scoped and do not conflict with identifiers that are not labels. The scope of a label is the body of the function in which it is declared and excludes the body of any nested function.

### Blank identifier ### {#Blank_identifier}

The _blank identifier_ is represented by the underscore character `_`. It serves as an anonymous placeholder instead of a regular (non-blank) identifier and has special meaning in [declarations](#Declarations_and_scope), as an [operand](#Operands), and in [assignments](#Assignments).

### Predeclared identifiers ### {#Predeclared_identifiers}

The following identifiers are implicitly declared in the [universe block](#Blocks):

``` go
Types:
	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap close complex copy delete imag len
	make new panic print println real recover
```

### Exported identifiers ### {#Exported_identifiers}

An identifier may be _exported_ to permit access to it from another package. An identifier is exported if both:

1.  the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and
2.  the identifier is declared in the [package block](#Blocks) or it is a [field name](#Struct_types) or [method name](#MethodName).

All other identifiers are not exported.

### Uniqueness of identifiers ### {#Uniqueness_of_identifiers}

Given a set of identifiers, an identifier is called _unique_ if it is _different_ from every other in the set. Two identifiers are different if they are spelled differently, or if they appear in different [packages](#Packages) and are not [exported](#Exported_identifiers). Otherwise, they are the same.

### Constant declarations ### {#Constant_declarations}

A constant declaration binds a list of identifiers (the names of the constants) to the values of a list of [constant expressions](#Constant_expressions). The number of identifiers must be equal to the number of expressions, and the _n_th identifier on the left is bound to the value of the _n_th expression on the right.

<pre class="ebnf"><a id="ConstDecl">ConstDecl</a>      = "const" ( <a href="#ConstSpec" class="noline">ConstSpec</a> | "(" { <a href="#ConstSpec" class="noline">ConstSpec</a> ";" } ")" ) .
<a id="ConstSpec">ConstSpec</a>      = <a href="#IdentifierList" class="noline">IdentifierList</a> [ [ <a href="#Type" class="noline">Type</a> ] "=" <a href="#ExpressionList" class="noline">ExpressionList</a> ] .

<a id="IdentifierList">IdentifierList</a> = <a href="#identifier" class="noline">identifier</a> { "," <a href="#identifier" class="noline">identifier</a> } .
<a id="ExpressionList">ExpressionList</a> = <a href="#Expression" class="noline">Expression</a> { "," <a href="#Expression" class="noline">Expression</a> } .
</pre>

If the type is present, all constants take the type specified, and the expressions must be [assignable](#Assignability) to that type. If the type is omitted, the constants take the individual types of the corresponding expressions. If the expression values are untyped [constants](#Constants), the declared constants remain untyped and the constant identifiers denote the constant values. For instance, if the expression is a floating-point literal, the constant identifier denotes a floating-point constant, even if the literal's fractional part is zero.

``` go
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1  // untyped integer constant
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
```

Within a parenthesized `const` declaration list the expression list may be omitted from any but the first declaration. Such an empty list is equivalent to the textual substitution of the first preceding non-empty expression list and its type if any. Omitting the list of expressions is therefore equivalent to repeating the previous list. The number of identifiers must be equal to the number of expressions in the previous list. Together with the [`iota` constant generator](#Iota) this mechanism permits light-weight declaration of sequential values:

``` go
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
```

### Iota ### {#Iota}

Within a [constant declaration](#Constant_declarations), the predeclared identifier `iota` represents successive untyped integer [constants](#Constants). It is reset to 0 whenever the reserved word `const` appears in the source and increments after each [ConstSpec](#ConstSpec). It can be used to construct a set of related constants:

``` go
const ( // iota is reset to 0
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const ( // iota is reset to 0
	a = 1 << iota  // a == 1
	b = 1 << iota  // b == 2
	c = 3          // c == 3  (iota is not used but still incremented)
	d = 1 << iota  // d == 8
)

const ( // iota is reset to 0
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0  (iota has been reset)
const y = iota  // y == 0  (iota has been reset)
```

Within an ExpressionList, the value of each `iota` is the same because it is only incremented after each ConstSpec:

``` go
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
	bit1, mask1                           // bit1 == 2, mask1 == 1
	_, _                                  // skips iota == 2
	bit3, mask3                           // bit3 == 8, mask3 == 7
)
```

This last example exploits the implicit repetition of the last non-empty expression list.

### Type declarations ### {#Type_declarations}

A type declaration binds an identifier, the _type name_, to a new type that has the same [underlying type](#Types) as an existing type, and operations defined for the existing type are also defined for the new type. The new type is [different](#Type_identity) from the existing type.

<pre class="ebnf"><a id="TypeDecl">TypeDecl</a>     = "type" ( <a href="#TypeSpec" class="noline">TypeSpec</a> | "(" { <a href="#TypeSpec" class="noline">TypeSpec</a> ";" } ")" ) .
<a id="TypeSpec">TypeSpec</a>     = <a href="#identifier" class="noline">identifier</a> <a href="#Type" class="noline">Type</a> .
</pre>

``` go
type IntArray [16]int

type (
	Point struct{ x, y float64 }
	Polar Point
)

type TreeNode struct {
	left, right *TreeNode
	value *Comparable
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
```

The declared type does not inherit any [methods](#Method_declarations) bound to the existing type, but the [method set](#Method_sets) of an interface type or of elements of a composite type remains unchanged:

``` go
// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of the [base type](#Pointer_types) of PtrMutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its anonymous field Mutex.
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block
```

A type declaration may be used to define a different boolean, numeric, or string type and attach methods to it:

``` go
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}
```

### Variable declarations ### {#Variable_declarations}

A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value.

<pre class="ebnf"><a id="VarDecl">VarDecl</a>     = "var" ( <a href="#VarSpec" class="noline">VarSpec</a> | "(" { <a href="#VarSpec" class="noline">VarSpec</a> ";" } ")" ) .
<a id="VarSpec">VarSpec</a>     = <a href="#IdentifierList" class="noline">IdentifierList</a> ( <a href="#Type" class="noline">Type</a> [ "=" <a href="#ExpressionList" class="noline">ExpressionList</a> ] | "=" <a href="#ExpressionList" class="noline">ExpressionList</a> ) .
</pre>

``` go
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
```

If a list of expressions is given, the variables are initialized with the expressions following the rules for [assignments](#Assignments). Otherwise, each variable is initialized to its [zero value](#The_zero_value).

If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment. If that value is an untyped constant, it is first [converted](#Conversions) to its [default type](#Constants); if it is an untyped boolean value, it is first converted to type `bool`. The predeclared value `nil` cannot be used to initialize a variable with no explicit type.

``` go
var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal
```

Implementation restriction: A compiler may make it illegal to declare a variable inside a [function body](#Function_declarations) if the variable is never used.

### Short variable declarations ### {#Short_variable_declarations}

A _short variable declaration_ uses the syntax:

<pre class="ebnf"><a id="ShortVarDecl">ShortVarDecl</a> = <a href="#IdentifierList" class="noline">IdentifierList</a> ":=" <a href="#ExpressionList" class="noline">ExpressionList</a> .
</pre>

It is shorthand for a regular [variable declaration](#Variable_declarations) with initializer expressions but no types:

``` go
"var" IdentifierList = ExpressionList .
```

``` go
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w := os.Pipe(fd)  // os.Pipe() returns two values
_, y, _ := coord(p)  // coord() returns three values; only interested in y coordinate
```

Unlike regular variable declarations, a short variable declaration may _redeclare_ variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-[blank](#Blank_identifier) variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original.

``` go
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
a, a := 1, 2                              // illegal: double declaration of a or no new variable if a was declared elsewhere
```

Short variable declarations may appear only inside functions. In some contexts such as the initializers for ["if"](#If_statements), ["for"](#For_statements), or ["switch"](#Switch_statements) statements, they can be used to declare local temporary variables.

### Function declarations ### {#Function_declarations}

A function declaration binds an identifier, the _function name_, to a function.

<pre class="ebnf"><a id="FunctionDecl">FunctionDecl</a> = "func" <a href="#FunctionName" class="noline">FunctionName</a> ( <a href="#Function" class="noline">Function</a> | <a href="#Signature" class="noline">Signature</a> ) .
<a id="FunctionName">FunctionName</a> = <a href="#identifier" class="noline">identifier</a> .
<a id="Function">Function</a>     = <a href="#Signature" class="noline">Signature</a> <a href="#FunctionBody" class="noline">FunctionBody</a> .
<a id="FunctionBody">FunctionBody</a> = <a href="#Block" class="noline">Block</a> .
</pre>

If the function's [signature](#Function_types) declares result parameters, the function body's statement list must end in a [terminating statement](#Terminating_statements).

``` go
func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	// invalid: missing return statement
}
```

A function declaration may omit the body. Such a declaration provides the signature for a function implemented outside Go, such as an assembly routine.

``` go
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func flushICache(begin, end uintptr)  // implemented externally
```

### Method declarations ### {#Method_declarations}

A method is a [function](#Function_declarations) with a _receiver_. A method declaration binds an identifier, the _method name_, to a method, and associates the method with the receiver's _base type_.

<pre class="ebnf"><a id="MethodDecl">MethodDecl</a>   = "func" <a href="#Receiver" class="noline">Receiver</a> <a href="#MethodName" class="noline">MethodName</a> ( <a href="#Function" class="noline">Function</a> | <a href="#Signature" class="noline">Signature</a> ) .
<a id="Receiver">Receiver</a>     = <a href="#Parameters" class="noline">Parameters</a> .
</pre>

The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be of the form `T` or `*T` (possibly using parentheses) where `T` is a type name. The type denoted by `T` is called the receiver _base type_; it must not be a pointer or interface type and it must be declared in the same package as the method. The method is said to be _bound_ to the base type and the method name is visible only within [selectors](#Selectors) for type `T` or `*T`.

A non-[blank](#Blank_identifier) receiver identifier must be [unique](#Uniqueness_of_identifiers) in the method signature. If the receiver's value is not referenced inside the body of the method, its identifier may be omitted in the declaration. The same applies in general to parameters of functions and methods.

For a base type, the non-blank names of methods bound to it must be unique. If the base type is a [struct type](#Struct_types), the non-blank method and field names must be distinct.

Given type `Point`, the declarations

``` go
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```

bind the methods `Length` and `Scale`, with receiver type `*Point`, to the base type `Point`.

The type of a method is the type of a function with the receiver as first argument. For instance, the method `Scale` has type

``` go
func(p *Point, factor float64)
```

However, a function declared this way is not a method.
