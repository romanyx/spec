## Expressions ## {#Expressions}

An expression specifies the computation of a value by applying operators and functions to operands.

### Operands ### {#Operands}

Operands denote the elementary values in an expression. An operand may be a literal, a (possibly [qualified](#Qualified_identifiers)) non-[blank](#Blank_identifier) identifier denoting a [constant](#Constant_declarations), [variable](#Variable_declarations), or [function](#Function_declarations), a [method expression](#Method_expressions) yielding a function, or a parenthesized expression.

The [blank identifier](#Blank_identifier) may appear as an operand only on the left-hand side of an [assignment](#Assignments).

<pre class="ebnf"><a id="Operand">Operand</a>     = <a href="#Literal" class="noline">Literal</a> | <a href="#OperandName" class="noline">OperandName</a> | <a href="#MethodExpr" class="noline">MethodExpr</a> | "(" <a href="#Expression" class="noline">Expression</a> ")" .
<a id="Literal">Literal</a>     = <a href="#BasicLit" class="noline">BasicLit</a> | <a href="#CompositeLit" class="noline">CompositeLit</a> | <a href="#FunctionLit" class="noline">FunctionLit</a> .
<a id="BasicLit">BasicLit</a>    = <a href="#int_lit" class="noline">int_lit</a> | <a href="#float_lit" class="noline">float_lit</a> | <a href="#imaginary_lit" class="noline">imaginary_lit</a> | <a href="#rune_lit" class="noline">rune_lit</a> | <a href="#string_lit" class="noline">string_lit</a> .
<a id="OperandName">OperandName</a> = <a href="#identifier" class="noline">identifier</a> | <a href="#QualifiedIdent" class="noline">QualifiedIdent</a>.
</pre>

### Qualified identifiers ### {#Qualified_identifiers}

A qualified identifier is an identifier qualified with a package name prefix. Both the package name and the identifier must not be [blank](#Blank_identifier).

<pre class="ebnf"><a id="QualifiedIdent">QualifiedIdent</a> = <a href="#PackageName" class="noline">PackageName</a> "." <a href="#identifier" class="noline">identifier</a> .
</pre>

A qualified identifier accesses an identifier in a different package, which must be [imported](#Import_declarations). The identifier must be [exported](#Exported_identifiers) and declared in the [package block](#Blocks) of that package.

``` go
math.Sin	// denotes the Sin function in package math
```

### Composite literals ### {#Composite_literals}

Composite literals construct values for structs, arrays, slices, and maps and create a new value each time they are evaluated. They consist of the type of the literal followed by a brace-bound list of elements. Each element may optionally be preceded by a corresponding key.

<pre class="ebnf"><a id="CompositeLit">CompositeLit</a>  = <a href="#LiteralType" class="noline">LiteralType</a> <a href="#LiteralValue" class="noline">LiteralValue</a> .
<a id="LiteralType">LiteralType</a>   = <a href="#StructType" class="noline">StructType</a> | <a href="#ArrayType" class="noline">ArrayType</a> | "[" "..." "]" <a href="#ElementType" class="noline">ElementType</a> |
                <a href="#SliceType" class="noline">SliceType</a> | <a href="#MapType" class="noline">MapType</a> | <a href="#TypeName" class="noline">TypeName</a> .
<a id="LiteralValue">LiteralValue</a>  = "{" [ <a href="#ElementList" class="noline">ElementList</a> [ "," ] ] "}" .
<a id="ElementList">ElementList</a>   = <a href="#KeyedElement" class="noline">KeyedElement</a> { "," <a href="#KeyedElement" class="noline">KeyedElement</a> } .
<a id="KeyedElement">KeyedElement</a>  = [ <a href="#Key" class="noline">Key</a> ":" ] <a href="#Element" class="noline">Element</a> .
<a id="Key">Key</a>           = <a href="#FieldName" class="noline">FieldName</a> | <a href="#Expression" class="noline">Expression</a> | <a href="#LiteralValue" class="noline">LiteralValue</a> .
<a id="FieldName">FieldName</a>     = <a href="#identifier" class="noline">identifier</a> .
<a id="Element">Element</a>       = <a href="#Expression" class="noline">Expression</a> | <a href="#LiteralValue" class="noline">LiteralValue</a> .
</pre>

The LiteralType's underlying type must be a struct, array, slice, or map type (the grammar enforces this constraint except when the type is given as a TypeName). The types of the elements and keys must be [assignable](#Assignability) to the respective field, element, and key types of the literal type; there is no additional conversion. The key is interpreted as a field name for struct literals, an index for array and slice literals, and a key for map literals. For map literals, all elements must have a key. It is an error to specify multiple elements with the same field name or constant key value.

For struct literals the following rules apply:

*   A key must be a field name declared in the struct type.
*   An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
*   If any element has a key, every element must have a key.
*   An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
*   A literal may omit the element list; such a literal evaluates to the zero value for its type.
*   It is an error to specify an element for a non-exported field of a struct belonging to a different package.

Given the declarations

``` go
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }
```

one may write

``` go
origin := Point3D{}                            // zero value for Point3D
line := Line{origin, Point3D{y: -4, z: 12.3}}  // zero value for line.q.x
```

For array and slice literals the following rules apply:

*   Each element has an associated integer index marking its position in the array.
*   An element with a key uses the key as its index. The key must be a non-negative constant representable by a value of type `int`; and if it is typed it must be of integer type.
*   An element without a key uses the previous element's index plus one. If the first element has no key, its index is zero.

[Taking the address](#Address_operators) of a composite literal generates a pointer to a unique [variable](#Variables) initialized with the literal's value.

``` go
var pointer *Point3D = &Point3D{y: 1000}
```

The length of an array literal is the length specified in the literal type. If fewer elements than the length are provided in the literal, the missing elements are set to the zero value for the array element type. It is an error to provide elements with index values outside the index range of the array. The notation `...` specifies an array length equal to the maximum element index plus one.

``` go
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2
```

A slice literal describes the entire underlying array literal. Thus the length and capacity of a slice literal are the maximum element index plus one. A slice literal has the form

``` go
[]T{x1, x2, … xn}
```

and is shorthand for a slice operation applied to an array:

``` go
tmp := [n]T{x1, x2, … xn}
tmp[0 : n]
```

Within a composite literal of array, slice, or map type `T`, elements or map keys that are themselves composite literals may elide the respective literal type if it is identical to the element or key type of `T`. Similarly, elements or keys that are addresses of composite literals may elide the `&T` when the element or key type is `*T`.

``` go
[...]Point{{1.5, -3.5}, {0, 0}}     // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
[][]int{{1, 2, 3}, {4, 5}}          // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}
[][]Point{{{0, 1}, {1, 2}}}         // same as [][]Point{[]Point{Point{0, 1}, Point{1, 2}}}
map[string]Point{"orig": {0, 0}}    // same as map[string]Point{"orig": Point{0, 0}}
map[Point]string{{0, 0}: "orig"}    // same as map[Point]string{Point{0, 0}: "orig"}

type PPoint *Point
[2]*Point{{1.5, -3.5}, {}}          // same as [2]*Point{&Point{1.5, -3.5}, &Point{}}
[2]PPoint{{1.5, -3.5}, {}}          // same as [2]PPoint{PPoint(&Point{1.5, -3.5}), PPoint(&Point{})}
```

A parsing ambiguity arises when a composite literal using the TypeName form of the LiteralType appears as an operand between the [keyword](#Keywords) and the opening brace of the block of an "if", "for", or "switch" statement, and the composite literal is not enclosed in parentheses, square brackets, or curly braces. In this rare case, the opening brace of the literal is erroneously parsed as the one introducing the block of statements. To resolve the ambiguity, the composite literal must appear within parentheses.

``` go
if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }
```

Examples of valid array, slice, and map literals:

``` go
// list of prime numbers
primes := []int{2, 3, 5, 7, 9, 2147483647}

// vowels[ch] is true if ch is a vowel
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
```

### Function literals ### {#Function_literals}

A function literal represents an anonymous [function](#Function_declarations).

<pre class="ebnf"><a id="FunctionLit">FunctionLit</a> = "func" <a href="#Function" class="noline">Function</a> .
</pre>

``` go
func(a, b int, z float64) bool { return a*b < int(z) }
```

A function literal can be assigned to a variable or invoked directly.

``` go
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

Function literals are _closures_: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible.

### Primary expressions ### {#Primary_expressions}

Primary expressions are the operands for unary and binary expressions.

<pre class="ebnf"><a id="PrimaryExpr">PrimaryExpr</a> =
	<a href="#Operand" class="noline">Operand</a> |
	<a href="#Conversion" class="noline">Conversion</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Selector" class="noline">Selector</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Index" class="noline">Index</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Slice" class="noline">Slice</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#TypeAssertion" class="noline">TypeAssertion</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Arguments" class="noline">Arguments</a> .

<a id="Selector">Selector</a>       = "." <a href="#identifier" class="noline">identifier</a> .
<a id="Index">Index</a>          = "[" <a href="#Expression" class="noline">Expression</a> "]" .
<a id="Slice">Slice</a>          = "[" [ <a href="#Expression" class="noline">Expression</a> ] ":" [ <a href="#Expression" class="noline">Expression</a> ] "]" |
                 "[" [ <a href="#Expression" class="noline">Expression</a> ] ":" <a href="#Expression" class="noline">Expression</a> ":" <a href="#Expression" class="noline">Expression</a> "]" .
<a id="TypeAssertion">TypeAssertion</a>  = "." "(" <a href="#Type" class="noline">Type</a> ")" .
<a id="Arguments">Arguments</a>      = "(" [ ( <a href="#ExpressionList" class="noline">ExpressionList</a> | <a href="#Type" class="noline">Type</a> [ "," <a href="#ExpressionList" class="noline">ExpressionList</a> ] ) [ "..." ] [ "," ] ] ")" .
</pre>

``` go
x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()
```

### Selectors ### {#Selectors}

For a [primary expression](#Primary_expressions) `x` that is not a [package name](#Package_clause), the _selector expression_

``` go
x.f
```

denotes the field or method `f` of the value `x` (or sometimes `*x`; see below). The identifier `f` is called the (field or method) _selector_; it must not be the [blank identifier](#Blank_identifier). The type of the selector expression is the type of `f`. If `x` is a package name, see the section on [qualified identifiers](#Qualified_identifiers).

A selector `f` may denote a field or method `f` of a type `T`, or it may refer to a field or method `f` of a nested [anonymous field](#Struct_types) of `T`. The number of anonymous fields traversed to reach `f` is called its _depth_ in `T`. The depth of a field or method `f` declared in `T` is zero. The depth of a field or method `f` declared in an anonymous field `A` in `T` is the depth of `f` in `A` plus one.

The following rules apply to selectors:

1.  For a value `x` of type `T` or `*T` where `T` is not a pointer or interface type, `x.f` denotes the field or method at the shallowest depth in `T` where there is such an `f`. If there is not exactly [one `f`](#Uniqueness_of_identifiers) with shallowest depth, the selector expression is illegal.
2.  For a value `x` of type `I` where `I` is an interface type, `x.f` denotes the actual method with name `f` of the dynamic value of `x`. If there is no method with name `f` in the [method set](#Method_sets) of `I`, the selector expression is illegal.
3.  As an exception, if the type of `x` is a named pointer type and `(*x).f` is a valid selector expression denoting a field (but not a method), `x.f` is shorthand for `(*x).f`.
4.  In all other cases, `x.f` is illegal.
5.  If `x` is of pointer type and has the value `nil` and `x.f` denotes a struct field, assigning to or evaluating `x.f` causes a [run-time panic](#Run_time_panics).
6.  If `x` is of interface type and has the value `nil`, [calling](#Calls) or [evaluating](#Method_values) the method `x.f` causes a [run-time panic](#Run_time_panics).

For example, given the declarations:

``` go
type T0 struct {
	x int
}

func (*T0) M0()

type T1 struct {
	y int
}

func (T1) M1()

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2()

type Q *T2

var t T2     // with t.T0 != nil
var p *T2    // with p != nil and (*p).T0 != nil
var q Q = p
```

one may write:

``` go
t.z          // t.z
t.y          // t.T1.y
t.x          // (*t.T0).x

p.z          // (*p).z
p.y          // (*p).T1.y
p.x          // (*(*p).T0).x

q.x          // (*(*q).T0).x        (*q).x is a valid field selector

p.M0()       // ((*p).T0).M0()      M0 expects *T0 receiver
p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
p.M2()       // p.M2()              M2 expects *T2 receiver
t.M2()       // (&t).M2()           M2 expects *T2 receiver, see section on Calls
```

but the following is invalid:

``` go
q.M0()       // (*q).M0 is valid but not a field selector
```

### Method expressions ### {#Method_expressions}

If `M` is in the [method set](#Method_sets) of type `T`, `T.M` is a function that is callable as a regular function with the same arguments as `M` prefixed by an additional argument that is the receiver of the method.

<pre class="ebnf"><a id="MethodExpr">MethodExpr</a>    = <a href="#ReceiverType" class="noline">ReceiverType</a> "." <a href="#MethodName" class="noline">MethodName</a> .
<a id="ReceiverType">ReceiverType</a>  = <a href="#TypeName" class="noline">TypeName</a> | "(" "*" <a href="#TypeName" class="noline">TypeName</a> ")" | "(" <a href="#ReceiverType" class="noline">ReceiverType</a> ")" .
</pre>

Consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.

``` go
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
```

The expression

``` go
T.Mv
```

yields a function equivalent to `Mv` but with an explicit receiver as its first argument; it has signature

``` go
func(tv T, a int) int
```

That function may be called normally with an explicit receiver, so these five invocations are equivalent:

``` go
t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)
```

Similarly, the expression

``` go
(*T).Mp
```

yields a function value representing `Mp` with signature

``` go
func(tp *T, f float32) float32
```

For a method with a value receiver, one can derive a function with an explicit pointer receiver, so

``` go
(*T).Mv
```

yields a function value representing `Mv` with signature

``` go
func(tv *T, a int) int
```

Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call.

The final case, a value-receiver function for a pointer-receiver method, is illegal because pointer-receiver methods are not in the method set of the value type.

Function values derived from methods are called with function call syntax; the receiver is provided as the first argument to the call. That is, given `f := T.Mv`, `f` is invoked as `f(t, 7)` not `t.f(7)`. To construct a function that binds the receiver, use a [function literal](#Function_literals) or [method value](#Method_values).

It is legal to derive a function value from a method of an interface type. The resulting function takes an explicit receiver of that interface type.

### Method values ### {#Method_values}

If the expression `x` has static type `T` and `M` is in the [method set](#Method_sets) of type `T`, `x.M` is called a _method value_. The method value `x.M` is a function value that is callable with the same arguments as a method call of `x.M`. The expression `x` is evaluated and saved during the evaluation of the method value; the saved copy is then used as the receiver in any calls, which may be executed later.

The type `T` may be an interface or non-interface type.

As in the discussion of [method expressions](#Method_expressions) above, consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.

``` go
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
var pt *T
func makeT() T
```

The expression

``` go
t.Mv
```

yields a function value of type

``` go
func(int) int
```

These two invocations are equivalent:

``` go
t.Mv(7)
f := t.Mv; f(7)
```

Similarly, the expression

``` go
pt.Mp
```

yields a function value of type

``` go
func(float32) float32
```

As with [selectors](#Selectors), a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`.

As with [method calls](#Calls), a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`.

``` go
f := t.Mv; f(7)   // like t.Mv(7)
f := pt.Mp; f(7)  // like pt.Mp(7)
f := pt.Mv; f(7)  // like (*pt).Mv(7)
f := t.Mp; f(7)   // like (&t).Mp(7)
f := makeT().Mp   // invalid: result of makeT() is not addressable
```

Although the examples above use non-interface types, it is also legal to create a method value from a value of interface type.

``` go
var i interface { M(int) } = myVal
f := i.M; f(7)  // like i.M(7)
```

### Index expressions ### {#Index_expressions}

A primary expression of the form

``` go
a[x]
```

denotes the element of the array, pointer to array, slice, string or map `a` indexed by `x`. The value `x` is called the _index_ or _map key_, respectively. The following rules apply:

If `a` is not a map:

*   the index `x` must be of integer type or untyped; it is _in range_ if `0 <= x < len(a)`, otherwise it is _out of range_
*   a [constant](#Constants) index must be non-negative and representable by a value of type `int`

For `a` of [array type](#Array_types) `A`:

*   a [constant](#Constants) index must be in range
*   if `x` is out of range at run time, a [run-time panic](#Run_time_panics) occurs
*   `a[x]` is the array element at index `x` and the type of `a[x]` is the element type of `A`

For `a` of [pointer](#Pointer_types) to array type:

*   `a[x]` is shorthand for `(*a)[x]`

For `a` of [slice type](#Slice_types) `S`:

*   if `x` is out of range at run time, a [run-time panic](#Run_time_panics) occurs
*   `a[x]` is the slice element at index `x` and the type of `a[x]` is the element type of `S`

For `a` of [string type](#String_types):

*   a [constant](#Constants) index must be in range if the string `a` is also constant
*   if `x` is out of range at run time, a [run-time panic](#Run_time_panics) occurs
*   `a[x]` is the non-constant byte value at index `x` and the type of `a[x]` is `byte`
*   `a[x]` may not be assigned to

For `a` of [map type](#Map_types) `M`:

*   `x`'s type must be [assignable](#Assignability) to the key type of `M`
*   if the map contains an entry with key `x`, `a[x]` is the map value with key `x` and the type of `a[x]` is the value type of `M`
*   if the map is `nil` or does not contain such an entry, `a[x]` is the [zero value](#The_zero_value) for the value type of `M`

Otherwise `a[x]` is illegal.

An index expression on a map `a` of type `map[K]V` used in an [assignment](#Assignments) or initialization of the special form

``` go
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
var v, ok T = a[x]
```

yields an additional untyped boolean value. The value of `ok` is `true` if the key `x` is present in the map, and `false` otherwise.

Assigning to an element of a `nil` map causes a [run-time panic](#Run_time_panics).

### Slice expressions ### {#Slice_expressions}

Slice expressions construct a substring or slice from a string, array, pointer to array, or slice. There are two variants: a simple form that specifies a low and high bound, and a full form that also specifies a bound on the capacity.

#### Simple slice expressions

For a string, array, pointer to array, or slice `a`, the primary expression

``` go
a[low : high]
```

constructs a substring or slice. The _indices_ `low` and `high` select which elements of operand `a` appear in the result. The result has indices starting at 0 and length equal to `high` - `low`. After slicing the array `a`

``` go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```

the slice `s` has type `[]int`, length 3, capacity 4, and elements

``` go
s[0] == 2
s[1] == 3
s[2] == 4
```

For convenience, any of the indices may be omitted. A missing `low` index defaults to zero; a missing `high` index defaults to the length of the sliced operand:

``` go
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
```

If `a` is a pointer to an array, `a[low : high]` is shorthand for `(*a)[low : high]`.

For arrays or strings, the indices are _in range_ if `0` <= `low` <= `high` <= `len(a)`, otherwise they are _out of range_. For slices, the upper index bound is the slice capacity `cap(a)` rather than the length. A [constant](#Constants) index must be non-negative and representable by a value of type `int`; for arrays or constant strings, constant indices must also be in range. If both indices are constant, they must satisfy `low <= high`. If the indices are out of range at run time, a [run-time panic](#Run_time_panics) occurs.

Except for [untyped strings](#Constants), if the sliced operand is a string or slice, the result of the slice operation is a non-constant value of the same type as the operand. For untyped string operands the result is a non-constant value of type `string`. If the sliced operand is an array, it must be [addressable](#Address_operators) and the result of the slice operation is a slice with the same element type as the array.

If the sliced operand of a valid slice expression is a `nil` slice, the result is a `nil` slice. Otherwise, the result shares its underlying array with the operand.

#### Full slice expressions

For an array, pointer to array, or slice `a` (but not a string), the primary expression

``` go
a[low : high : max]
```

constructs a slice of the same type, and with the same length and elements as the simple slice expression `a[low : high]`. Additionally, it controls the resulting slice's capacity by setting it to `max - low`. Only the first index may be omitted; it defaults to 0. After slicing the array `a`

``` go
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
```

the slice `t` has type `[]int`, length 2, capacity 4, and elements

``` go
t[0] == 2
t[1] == 3
```

As for simple slice expressions, if `a` is a pointer to an array, `a[low : high : max]` is shorthand for `(*a)[low : high : max]`. If the sliced operand is an array, it must be [addressable](#Address_operators).

The indices are _in range_ if `0 <= low <= high <= max <= cap(a)`, otherwise they are _out of range_. A [constant](#Constants) index must be non-negative and representable by a value of type `int`; for arrays, constant indices must also be in range. If multiple indices are constant, the constants that are present must be in range relative to each other. If the indices are out of range at run time, a [run-time panic](#Run_time_panics) occurs.

### Type assertions ### {#Type_assertions}

For an expression `x` of [interface type](#Interface_types) and a type `T`, the primary expression

``` go
x.(T)
```

asserts that `x` is not `nil` and that the value stored in `x` is of type `T`. The notation `x.(T)` is called a _type assertion_.

More precisely, if `T` is not an interface type, `x.(T)` asserts that the dynamic type of `x` is [identical](#Type_identity) to the type `T`. In this case, `T` must [implement](#Method_sets) the (interface) type of `x`; otherwise the type assertion is invalid since it is not possible for `x` to store a value of type `T`. If `T` is an interface type, `x.(T)` asserts that the dynamic type of `x` implements the interface `T`.

If the type assertion holds, the value of the expression is the value stored in `x` and its type is `T`. If the type assertion is false, a [run-time panic](#Run_time_panics) occurs. In other words, even though the dynamic type of `x` is known only at run time, the type of `x.(T)` is known to be `T` in a correct program.

``` go
var x interface{} = 7          // x has dynamic type int and value 7
i := x.(int)                   // i has type int and value 7

type I interface { m() }

func f(y I) {
	s := y.(string)        // illegal: string does not implement I (missing method m)
	r := y.(io.Reader)     // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
	…
}
```

A type assertion used in an [assignment](#Assignments) or initialization of the special form

``` go
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
var v, ok T1 = x.(T)
```

yields an additional untyped boolean value. The value of `ok` is `true` if the assertion holds. Otherwise it is `false` and the value of `v` is the [zero value](#The_zero_value) for type `T`. No run-time panic occurs in this case.

### Calls ### {#Calls}

Given an expression `f` of function type `F`,

``` go
f(a1, a2, … an)
```

calls `f` with arguments `a1, a2, … an`. Except for one special case, arguments must be single-valued expressions [assignable](#Assignability) to the parameter types of `F` and are evaluated before the function is called. The type of the expression is the result type of `F`. A method invocation is similar but the method itself is specified as a selector upon a value of the receiver type for the method.

``` go
math.Atan2(x, y)  // function call
var pt *Point
pt.Scale(3.5)     // method call with receiver pt
```

In a function call, the function value and arguments are evaluated in [the usual order](#Order_of_evaluation). After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution. The return parameters of the function are passed by value back to the calling function when the function returns.

Calling a `nil` function value causes a [run-time panic](#Run_time_panics).

As a special case, if the return values of a function or method `g` are equal in number and individually assignable to the parameters of another function or method `f`, then the call `f(g(_parameters_of_g_))` will invoke `f` after binding the return values of `g` to the parameters of `f` in order. The call of `f` must contain no parameters other than the call of `g`, and `g` must have at least one return value. If `f` has a final `...` parameter, it is assigned the return values of `g` that remain after assignment of regular parameters.

``` go
func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}
```

A method call `x.m()` is valid if the [method set](#Method_sets) of (the type of) `x` contains `m` and the argument list can be assigned to the parameter list of `m`. If `x` is [addressable](#Address_operators) and `&x`'s method set contains `m`, `x.m()` is shorthand for `(&x).m()`:

``` go
var p Point
p.Scale(3.5)
```

There is no distinct method type and there are no method literals.

### Passing arguments to `...` parameters ### {#Passing_arguments_to_..._parameters}

If `f` is [variadic](#Function_types) with a final parameter `p` of type `...T`, then within `f` the type of `p` is equivalent to type `[]T`. If `f` is invoked with no actual arguments for `p`, the value passed to `p` is `nil`. Otherwise, the value passed is a new slice of type `[]T` with a new underlying array whose successive elements are the actual arguments, which all must be [assignable](#Assignability) to `T`. The length and capacity of the slice is therefore the number of arguments bound to `p` and may differ for each call site.

Given the function and calls

``` go
func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
```

within `Greeting`, `who` will have the value `nil` in the first call, and `[]string{"Joe", "Anna", "Eileen"}` in the second.

If the final argument is assignable to a slice type `[]T`, it may be passed unchanged as the value for a `...T` parameter if the argument is followed by `...`. In this case no new slice is created.

Given the slice `s` and call

``` go
s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
```

within `Greeting`, `who` will have the same value as `s` with the same underlying array.

### Operators ### {#Operators}

Operators combine operands into expressions.

<pre class="ebnf"><a id="Expression">Expression</a> = <a href="#UnaryExpr" class="noline">UnaryExpr</a> | <a href="#Expression" class="noline">Expression</a> <a href="#binary_op" class="noline">binary_op</a> <a href="#Expression" class="noline">Expression</a> .
<a id="UnaryExpr">UnaryExpr</a>  = <a href="#PrimaryExpr" class="noline">PrimaryExpr</a> | <a href="#unary_op" class="noline">unary_op</a> <a href="#UnaryExpr" class="noline">UnaryExpr</a> .

<a id="binary_op">binary_op</a>  = "||" | "&amp;&amp;" | <a href="#rel_op" class="noline">rel_op</a> | <a href="#add_op" class="noline">add_op</a> | <a href="#mul_op" class="noline">mul_op</a> .
<a id="rel_op">rel_op</a>     = "==" | "!=" | "&lt;" | "&lt;=" | "&gt;" | "&gt;=" .
<a id="add_op">add_op</a>     = "+" | "-" | "|" | "^" .
<a id="mul_op">mul_op</a>     = "*" | "/" | "%" | "&lt;&lt;" | "&gt;&gt;" | "&amp;" | "&amp;^" .

<a id="unary_op">unary_op</a>   = "+" | "-" | "!" | "^" | "*" | "&amp;" | "&lt;-" .
</pre>

Comparisons are discussed [elsewhere](#Comparison_operators). For other binary operators, the operand types must be [identical](#Type_identity) unless the operation involves shifts or untyped [constants](#Constants). For operations involving constants only, see the section on [constant expressions](#Constant_expressions).

Except for shift operations, if one operand is an untyped [constant](#Constants) and the other operand is not, the constant is [converted](#Conversions) to the type of the other operand.

The right operand in a shift expression must have unsigned integer type or be an untyped constant that can be converted to unsigned integer type. If the left operand of a non-constant shift expression is an untyped constant, it is first converted to the type it would assume if the shift expression were replaced by its left operand alone.

``` go
var s uint = 33
var i = 1<<s           // 1 has type int
var j int32 = 1<<s     // 1 has type int32; j == 0
var k = uint64(1<<s)   // 1 has type uint64; k == 1<<33
var m int = 1.0<<s     // 1.0 has type int; m == 0 if ints are 32bits in size
var n = 1.0<<s == j    // 1.0 has type int32; n == true
var o = 1<<s == 2<<s   // 1 and 2 have type int; o == true if ints are 32bits in size
var p = 1<<s == 1<<33  // illegal if ints are 32bits in size: 1 has type int, but 1<<33 overflows int
var u = 1.0<<s         // illegal: 1.0 has type float64, cannot shift
var u1 = 1.0<<s != 0   // illegal: 1.0 has type float64, cannot shift
var u2 = 1<<s != 1.0   // illegal: 1 has type float64, cannot shift
var v float32 = 1<<s   // illegal: 1 has type float32, cannot shift
var w int64 = 1.0<<33  // 1.0<<33 is a constant shift expression
```

#### Operator precedence

Unary operators have the highest precedence. As the `++` and `--` operators form statements, not expressions, they fall outside the operator hierarchy. As a consequence, statement `*p++` is the same as `(*p)++`.

There are five precedence levels for binary operators. Multiplication operators bind strongest, followed by addition operators, comparison operators, `&&` (logical AND), and finally `||` (logical OR):

``` go
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

Binary operators of the same precedence associate from left to right. For instance, `x / y * z` is the same as `(x / y) * z`.

``` go
+x
23 + 3*x[i]
x <= f()
^a >> b
f() || g()
x == y+1 && <-chanPtr > 0
```

### Arithmetic operators ### {#Arithmetic_operators}

Arithmetic operators apply to numeric values and yield a result of the same type as the first operand. The four standard arithmetic operators (`+`, `-`, `*`, `/`) apply to integer, floating-point, and complex types; `+` also applies to strings. The bitwise logical and shift operators apply to integers only.

``` go
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer
```

#### Integer operators

For two integer values `x` and `y`, the integer quotient `q = x / y` and remainder `r = x % y` satisfy the following relationships:

``` go
x = q*y + r  and  |r| < |y|
```

with `x / y` truncated towards zero (["truncated division"](http://en.wikipedia.org/wiki/Modulo_operation)).

``` go
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
```

As an exception to this rule, if the dividend `x` is the most negative value for the int type of `x`, the quotient `q = x / -1` is equal to `x` (and `r = 0`).

``` go
			 x, q
int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808
```

If the divisor is a [constant](#Constants), it must not be zero. If the divisor is zero at run time, a [run-time panic](#Run_time_panics) occurs. If the dividend is non-negative and the divisor is a constant power of 2, the division may be replaced by a right shift, and computing the remainder may be replaced by a bitwise AND operation:

``` go
 x     x / 4     x % 4     x >> 2     x & 3
 11      2         3         2          3
-11     -2        -3        -3          1
```

The shift operators shift the left operand by the shift count specified by the right operand. They implement arithmetic shifts if the left operand is a signed integer and logical shifts if it is an unsigned integer. There is no upper limit on the shift count. Shifts behave as if the left operand is shifted `n` times by 1 for a shift count of `n`. As a result, `x << 1` is the same as `x*2` and `x >> 1` is the same as `x/2` but truncated towards negative infinity.

For integer operands, the unary operators `+`, `-`, and `^` are defined as follows:

``` go
+x                          is 0 + x
-x    negation              is 0 - x
^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                      and  m = -1 for signed x
```

#### Integer overflow

For unsigned integer values, the operations `+`, `-`, `*`, and `<<` are computed modulo 2<sup>_n_</sup>, where _n_ is the bit width of the [unsigned integer](#Numeric_types)'s type. Loosely speaking, these unsigned integer operations discard high bits upon overflow, and programs may rely on ``wrap around''.

For signed integers, the operations `+`, `-`, `*`, and `<<` may legally overflow and the resulting value exists and is deterministically defined by the signed integer representation, the operation, and its operands. No exception is raised as a result of overflow. A compiler may not optimize code under the assumption that overflow does not occur. For instance, it may not assume that `x < x + 1` is always true.

#### Floating-point operators

For floating-point and complex numbers, `+x` is the same as `x`, while `-x` is the negation of `x`. The result of a floating-point or complex division by zero is not specified beyond the IEEE-754 standard; whether a [run-time panic](#Run_time_panics) occurs is implementation-specific.

#### String concatenation

Strings can be concatenated using the `+` operator or the `+=` assignment operator:

``` go
s := "hi" + string(c)
s += " and good bye"
```

String addition creates a new string by concatenating the operands.

### Comparison operators ### {#Comparison_operators}

Comparison operators compare two operands and yield an untyped boolean value.

``` go
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

In any comparison, the first operand must be [assignable](#Assignability) to the type of the second operand, or vice versa.

The equality operators `==` and `!=` apply to operands that are _comparable_. The ordering operators `<`, `<=`, `>`, and `>=` apply to operands that are _ordered_. These terms and the result of the comparisons are defined as follows:

*   Boolean values are comparable. Two boolean values are equal if they are either both `true` or both `false`.
*   Integer values are comparable and ordered, in the usual way.
*   Floating point values are comparable and ordered, as defined by the IEEE-754 standard.
*   Complex values are comparable. Two complex values `u` and `v` are equal if both `real(u) == real(v)` and `imag(u) == imag(v)`.
*   String values are comparable and ordered, lexically byte-wise.
*   Pointer values are comparable. Two pointer values are equal if they point to the same variable or if both have value `nil`. Pointers to distinct [zero-size](#Size_and_alignment_guarantees) variables may or may not be equal.
*   Channel values are comparable. Two channel values are equal if they were created by the same call to [`make`](#Making_slices_maps_and_channels) or if both have value `nil`.
*   Interface values are comparable. Two interface values are equal if they have [identical](#Type_identity) dynamic types and equal dynamic values or if both have value `nil`.
*   A value `x` of non-interface type `X` and a value `t` of interface type `T` are comparable when values of type `X` are comparable and `X` implements `T`. They are equal if `t`'s dynamic type is identical to `X` and `t`'s dynamic value is equal to `x`.
*   Struct values are comparable if all their fields are comparable. Two struct values are equal if their corresponding non-[blank](#Blank_identifier) fields are equal.
*   Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.

A comparison of two interface values with identical dynamic types causes a [run-time panic](#Run_time_panics) if values of that type are not comparable. This behavior applies not only to direct interface value comparisons but also when comparing arrays of interface values or structs with interface-valued fields.

Slice, map, and function values are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`. Comparison of pointer, channel, and interface values to `nil` is also allowed and follows from the general rules above.

``` go
const c = 3 < 4            // c is the untyped boolean constant true

type MyBool bool
var x, y int
var (
	// The result of a comparison is an untyped boolean.
	// The usual assignment rules apply.
	b3        = x == y // b3 has type bool
	b4 bool   = x == y // b4 has type bool
	b5 MyBool = x == y // b5 has type MyBool
)
```

### Logical operators ### {#Logical_operators}

Logical operators apply to [boolean](#Boolean_types) values and yield a result of the same type as the operands. The right operand is evaluated conditionally.

``` go
&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```

### Address operators ### {#Address_operators}

For an operand `x` of type `T`, the address operation `&x` generates a pointer of type `*T` to `x`. The operand must be _addressable_, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, `x` may also be a (possibly parenthesized) [composite literal](#Composite_literals). If the evaluation of `x` would cause a [run-time panic](#Run_time_panics), then the evaluation of `&x` does too.

For an operand `x` of pointer type `*T`, the pointer indirection `*x` denotes the [variable](#Variables) of type `T` pointed to by `x`. If `x` is `nil`, an attempt to evaluate `*x` will cause a [run-time panic](#Run_time_panics).

``` go
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

### Receive operator ### {#Receive_operator}

For an operand `ch` of [channel type](#Channel_types), the value of the receive operation `<-ch` is the value received from the channel `ch`. The channel direction must permit receive operations, and the type of the receive operation is the element type of the channel. The expression blocks until a value is available. Receiving from a `nil` channel blocks forever. A receive operation on a [closed](#Close) channel can always proceed immediately, yielding the element type's [zero value](#The_zero_value) after any previously sent values have been received.

``` go
v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe  // wait until clock pulse and discard received value
```

A receive expression used in an [assignment](#Assignments) or initialization of the special form

``` go
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
```

yields an additional untyped boolean result reporting whether the communication succeeded. The value of `ok` is `true` if the value received was delivered by a successful send operation to the channel, or `false` if it is a zero value generated because the channel is closed and empty.

### Conversions ### {#Conversions}

Conversions are expressions of the form `T(x)` where `T` is a type and `x` is an expression that can be converted to type `T`.

<pre class="ebnf"><a id="Conversion">Conversion</a> = <a href="#Type" class="noline">Type</a> "(" <a href="#Expression" class="noline">Expression</a> [ "," ] ")" .
</pre>

If the type starts with the operator `*` or `<-`, or if the type starts with the keyword `func` and has no result list, it must be parenthesized when necessary to avoid ambiguity:

``` go
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
<-chan int(c)    // same as <-(chan int(c))
(<-chan int)(c)  // c is converted to <-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
```

A [constant](#Constants) value `x` can be converted to type `T` in any of these cases:

*   `x` is representable by a value of type `T`.
*   `x` is a floating-point constant, `T` is a floating-point type, and `x` is representable by a value of type `T` after rounding using IEEE 754 round-to-even rules, but with an IEEE `-0.0` further rounded to an unsigned `0.0`. The constant `T(x)` is the rounded value.
*   `x` is an integer constant and `T` is a [string type](#String_types). The [same rule](#Conversions_to_and_from_a_string_type) as for non-constant `x` applies in this case.

Converting a constant yields a typed constant as result.

``` go
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
MyString("foo" + "bar")  // "foobar" of type MyString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
```

A non-constant value `x` can be converted to type `T` in any of these cases:

*   `x` is [assignable](#Assignability) to `T`.
*   ignoring struct tags (see below), `x`'s type and `T` have [identical](#Type_identity) [underlying types](#Types).
*   ignoring struct tags (see below), `x`'s type and `T` are unnamed pointer types and their pointer base types have identical underlying types.
*   `x`'s type and `T` are both integer or floating point types.
*   `x`'s type and `T` are both complex types.
*   `x` is an integer or a slice of bytes or runes and `T` is a string type.
*   `x` is a string and `T` is a slice of bytes or runes.

[Struct tags](#Struct_types) are ignored when comparing struct types for identity for the purpose of conversion:

``` go
type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data *struct {
	Name    string `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
}

var person = (*Person)(data)  // ignoring tags, the underlying types are identical
```

Specific rules apply to (non-constant) conversions between numeric types or to and from a string type. These conversions may change the representation of `x` and incur a run-time cost. All other conversions only change the type but not the representation of `x`.

There is no linguistic mechanism to convert between pointers and integers. The package [`unsafe`](#Package_unsafe) implements this functionality under restricted circumstances.

#### Conversions between numeric types

For the conversion of non-constant numeric values, the following rules apply:

1.  When converting between integer types, if the value is a signed integer, it is sign extended to implicit infinite precision; otherwise it is zero extended. It is then truncated to fit in the result type's size. For example, if `v := uint16(0x10F0)`, then `uint32(int8(v)) == 0xFFFFFFF0`. The conversion always yields a valid value; there is no indication of overflow.
2.  When converting a floating-point number to an integer, the fraction is discarded (truncation towards zero).
3.  When converting an integer or floating-point number to a floating-point type, or a complex number to another complex type, the result value is rounded to the precision specified by the destination type. For instance, the value of a variable `x` of type `float32` may be stored using additional precision beyond that of an IEEE-754 32-bit number, but float32(x) represents the result of rounding `x`'s value to 32-bit precision. Similarly, `x + 0.1` may use more than 32 bits of precision, but `float32(x + 0.1)` does not.

In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.

#### Conversions to and from a string type

1.  Converting a signed or unsigned integer value to a string type yields a string containing the UTF-8 representation of the integer. Values outside the range of valid Unicode code points are converted to `"\uFFFD"`.

    <pre>string('a')       // "a"
    string(-1)        // "\ufffd" == "\xef\xbf\xbd"
    string(0xf8)      // "\u00f8" == "ø" == "\xc3\xb8"
    type MyString string
    MyString(0x65e5)  // "\u65e5" == "日" == "\xe6\x97\xa5"
    </pre>

2.  Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.

    <pre>string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})   // "hellø"
    string([]byte{})                                     // ""
    string([]byte(nil))                                  // ""

    type MyBytes []byte
    string(MyBytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})  // "hellø"
    </pre>

3.  Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.

    <pre>string([]rune{0x767d, 0x9d6c, 0x7fd4})   // "\u767d\u9d6c\u7fd4" == "白鵬翔"
    string([]rune{})                         // ""
    string([]rune(nil))                      // ""

    type MyRunes []rune
    string(MyRunes{0x767d, 0x9d6c, 0x7fd4})  // "\u767d\u9d6c\u7fd4" == "白鵬翔"
    </pre>

4.  Converting a value of a string type to a slice of bytes type yields a slice whose successive elements are the bytes of the string.

    <pre>[]byte("hellø")   // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
    []byte("")        // []byte{}

    MyBytes("hellø")  // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
    </pre>

5.  Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.

    <pre>[]rune(MyString("白鵬翔"))  // []rune{0x767d, 0x9d6c, 0x7fd4}
    []rune("")                 // []rune{}

    MyRunes("白鵬翔")           // []rune{0x767d, 0x9d6c, 0x7fd4}
    </pre>

### Constant expressions ### {#Constant_expressions}

Constant expressions may contain only [constant](#Constants) operands and are evaluated at compile time.

Untyped boolean, numeric, and string constants may be used as operands wherever it is legal to use an operand of boolean, numeric, or string type, respectively. Except for shift operations, if the operands of a binary operation are different kinds of untyped constants, the operation and, for non-boolean operations, the result use the kind that appears later in this list: integer, rune, floating-point, complex. For example, an untyped integer constant divided by an untyped complex constant yields an untyped complex constant.

A constant [comparison](#Comparison_operators) always yields an untyped boolean constant. If the left operand of a constant [shift expression](#Operators) is an untyped constant, the result is an integer constant; otherwise it is a constant of the same type as the left operand, which must be of [integer type](#Numeric_types). Applying all other operators to untyped constants results in an untyped constant of the same kind (that is, a boolean, integer, floating-point, complex, or string constant).

``` go
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2\.     // Π == 1.5   (type float64, 3/2\. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
```

Applying the built-in function `complex` to untyped integer, rune, or floating-point constants yields an untyped complex constant.

``` go
const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant)
const iΘ = complex(0, Θ)   // iΘ == 1i     (type complex128)
```

Constant expressions are always evaluated exactly; intermediate values and the constants themselves may require precision significantly larger than supported by any predeclared type in the language. The following are legal declarations:

``` go
const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge >> 98  // Four == 4                                (type int8)
```

The divisor of a constant division or remainder operation must not be zero:

``` go
3.14 / 0.0   // illegal: division by zero
```

The values of _typed_ constants must always be accurately representable as values of the constant type. The following constant expressions are illegal:

``` go
uint(-1)     // -1 cannot be represented as a uint
int(3.14)    // 3.14 cannot be represented as an int
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
Four * 100   // product 400 cannot be represented as an int8 (type of Four)
```

The mask used by the unary bitwise complement operator `^` matches the rule for non-constants: the mask is all 1s for unsigned constants and -1 for signed and untyped constants.

``` go
^1         // untyped integer constant, equal to -2
uint8(^1)  // illegal: same as uint8(-2), -2 cannot be represented as a uint8
^uint8(1)  // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // same as int8(-2)
^int8(1)   // same as -1 ^ int8(1) = -2
```

Implementation restriction: A compiler may use rounding while computing untyped floating-point or complex constant expressions; see the implementation restriction in the section on [constants](#Constants). This rounding may cause a floating-point constant expression to be invalid in an integer context, even if it would be integral when calculated using infinite precision, and vice versa.

### Order of evaluation ### {#Order_of_evaluation}

At package level, [initialization dependencies](#Package_initialization) determine the evaluation order of individual initialization expressions in [variable declarations](#Variable_declarations). Otherwise, when evaluating the [operands](#Operands) of an expression, assignment, or [return statement](#Return_statements), all function calls, method calls, and communication operations are evaluated in lexical left-to-right order.

For example, in the (function-local) assignment

``` go
y[f()], ok = g(h(), i()+x[j()], <-c), k()
```

the function calls and communication happen in the order `f()`, `h()`, `i()`, `j()`, `<-c`, `g()`, and `k()`. However, the order of those events compared to the evaluation and indexing of `x` and the evaluation of `y` is not specified.

``` go
a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified
```

At package level, initialization dependencies override the left-to-right rule for individual initialization expressions, but not for operands within each expression:

``` go
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// functions u and v are independent of all other variables and functions
```

The function calls happen in the order `u()`, `sqr()`, `v()`, `f()`, `v()`, and `g()`.

Floating-point operations within a single expression are evaluated according to the associativity of the operators. Explicit parentheses affect the evaluation by overriding the default associativity. In the expression `x + (y + z)` the addition `y + z` is performed before adding `x`.
