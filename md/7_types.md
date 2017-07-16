## Types ## {#Types}

A type determines the set of values and operations specific to values of that type. Types may be _named_ or _unnamed_. Named types are specified by a (possibly [qualified](#Qualified_identifiers)) [_type name_](#Type_declarations); unnamed types are specified using a _type literal_, which composes a new type from existing types.

<pre class="ebnf"><a id="Type">Type</a>      = <a href="#TypeName" class="noline">TypeName</a> | <a href="#TypeLit" class="noline">TypeLit</a> | "(" <a href="#Type" class="noline">Type</a> ")" .
<a id="TypeName">TypeName</a>  = <a href="#identifier" class="noline">identifier</a> | <a href="#QualifiedIdent" class="noline">QualifiedIdent</a> .
<a id="TypeLit">TypeLit</a>   = <a href="#ArrayType" class="noline">ArrayType</a> | <a href="#StructType" class="noline">StructType</a> | <a href="#PointerType" class="noline">PointerType</a> | <a href="#FunctionType" class="noline">FunctionType</a> | <a href="#InterfaceType" class="noline">InterfaceType</a> |
	    <a href="#SliceType" class="noline">SliceType</a> | <a href="#MapType" class="noline">MapType</a> | <a href="#ChannelType" class="noline">ChannelType</a> .
</pre>

Named instances of the boolean, numeric, and string types are [predeclared](#Predeclared_identifiers). _Composite types_—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.

Each type `T` has an _underlying type_: If `T` is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is `T` itself. Otherwise, `T`'s underlying type is the underlying type of the type to which `T` refers in its [type declaration](#Type_declarations).

``` go
type T1 string
type T2 T1
type T3 []T1
type T4 T3
```

The underlying type of `string`, `T1`, and `T2` is `string`. The underlying type of `[]T1`, `T3`, and `T4` is `[]T1`.

### Method sets ### {#Method_sets}

A type may have a _method set_ associated with it. The method set of an [interface type](#Interface_types) is its interface. The method set of any other type `T` consists of all [methods](#Method_declarations) declared with receiver type `T`. The method set of the corresponding [pointer type](#Pointer_types) `*T` is the set of all methods declared with receiver `*T` or `T` (that is, it also contains the method set of `T`). Further rules apply to structs containing anonymous fields, as described in the section on [struct types](#Struct_types). Any other type has an empty method set. In a method set, each method must have a [unique](#Uniqueness_of_identifiers) non-[blank](#Blank_identifier) [method name](#MethodName).

The method set of a type determines the interfaces that the type [implements](#Interface_types) and the methods that can be [called](#Calls) using a receiver of that type.

### Boolean types ### {#Boolean_types}

A _boolean type_ represents the set of Boolean truth values denoted by the predeclared constants `true` and `false`. The predeclared boolean type is `bool`.

### Numeric types ### {#Numeric_types}

A _numeric type_ represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are:

``` go
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

The value of an _n_-bit integer is _n_ bits wide and represented using [two's complement arithmetic](http://en.wikipedia.org/wiki/Two's_complement).

There is also a set of predeclared numeric types with implementation-specific sizes:

``` go
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

To avoid portability issues all numeric types are distinct except `byte`, which is an alias for `uint8`, and `rune`, which is an alias for `int32`. Conversions are required when different numeric types are mixed in an expression or assignment. For instance, `int32` and `int` are not the same type even though they may have the same size on a particular architecture.

### String types ### {#String_types}

A _string type_ represents the set of string values. A string value is a (possibly empty) sequence of bytes. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is `string`.

The length of a string `s` (its size in bytes) can be discovered using the built-in function [`len`](#Length_and_capacity). The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer [indices](#Index_expressions) 0 through `len(s)-1`. It is illegal to take the address of such an element; if `s[i]` is the `i`'th byte of a string, `&s[i]` is invalid.

### Array types ### {#Array_types}

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length and is never negative.

<pre class="ebnf"><a id="ArrayType">ArrayType</a>   = "[" <a href="#ArrayLength" class="noline">ArrayLength</a> "]" <a href="#ElementType" class="noline">ElementType</a> .
<a id="ArrayLength">ArrayLength</a> = <a href="#Expression" class="noline">Expression</a> .
<a id="ElementType">ElementType</a> = <a href="#Type" class="noline">Type</a> .
</pre>

The length is part of the array's type; it must evaluate to a non-negative [constant](#Constants) representable by a value of type `int`. The length of array `a` can be discovered using the built-in function [`len`](#Length_and_capacity). The elements can be addressed by integer [indices](#Index_expressions) 0 through `len(a)-1`. Array types are always one-dimensional but may be composed to form multi-dimensional types.

``` go
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

### Slice types ### {#Slice_types}

A slice is a descriptor for a contiguous segment of an _underlying array_ and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The value of an uninitialized slice is `nil`.

<pre class="ebnf"><a id="SliceType">SliceType</a> = "[" "]" <a href="#ElementType" class="noline">ElementType</a> .
</pre>

Like arrays, slices are indexable and have a length. The length of a slice `s` can be discovered by the built-in function [`len`](#Length_and_capacity); unlike with arrays it may change during execution. The elements can be addressed by integer [indices](#Index_expressions) 0 through `len(s)-1`. The slice index of a given element may be less than the index of the same element in the underlying array.

A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.

The array underlying a slice may extend past the end of the slice. The _capacity_ is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by [_slicing_](#Slice_expressions) a new one from the original slice. The capacity of a slice `a` can be discovered using the built-in function [`cap(a)`](#Length_and_capacity).

A new, initialized slice value for a given element type `T` is made using the built-in function [`make`](#Making_slices_maps_and_channels), which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with `make` always allocates a new, hidden array to which the returned slice value refers. That is, executing

``` go
make([]T, length, capacity)
```

produces the same slice as allocating an array and [slicing](#Slice_expressions) it, so these two expressions are equivalent:

``` go
make([]int, 50, 100)
new([100]int)[0:50]
```

Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.

### Struct types ### {#Struct_types}

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField). Within a struct, non-[blank](#Blank_identifier) field names must be [unique](#Uniqueness_of_identifiers).

<pre class="ebnf"><a id="StructType">StructType</a>     = "struct" "{" { <a href="#FieldDecl" class="noline">FieldDecl</a> ";" } "}" .
<a id="FieldDecl">FieldDecl</a>      = (<a href="#IdentifierList" class="noline">IdentifierList</a> <a href="#Type" class="noline">Type</a> | <a href="#AnonymousField" class="noline">AnonymousField</a>) [ <a href="#Tag" class="noline">Tag</a> ] .
<a id="AnonymousField">AnonymousField</a> = [ "*" ] <a href="#TypeName" class="noline">TypeName</a> .
<a id="Tag">Tag</a>            = <a href="#string_lit" class="noline">string_lit</a> .
</pre>

``` go
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

A field declared with a type but no explicit field name is an _anonymous field_, also called an _embedded_ field or an embedding of the type in the struct. An embedded type must be specified as a type name `T` or as a pointer to a non-interface type name `*T`, and `T` itself may not be a pointer type. The unqualified type name acts as the field name.

``` go 
// A struct with four anonymous fields of type T1, *T2, P.T3 and *P.T4
struct {
	T1        // field name is T1
	*T2       // field name is T2
	P.T3      // field name is T3
	*P.T4     // field name is T4
	x, y int  // field names are x and y
}
```

The following declaration is illegal because field names must be unique in a struct type:

``` go
struct {
	T     // conflicts with anonymous field *T and *P.T
	*T    // conflicts with anonymous field T and *P.T
	*P.T  // conflicts with anonymous field T and *T
}
```

A field or [method](#Method_declarations) `f` of an anonymous field in a struct `x` is called _promoted_ if `x.f` is a legal [selector](#Selectors) that denotes that field or method `f`.

Promoted fields act like ordinary fields of a struct except that they cannot be used as field names in [composite literals](#Composite_literals) of the struct.

Given a struct type `S` and a type named `T`, promoted methods are included in the method set of the struct as follows:

*   If `S` contains an anonymous field `T`, the [method sets](#Method_sets) of `S` and `*S` both include promoted methods with receiver `T`. The method set of `*S` also includes promoted methods with receiver `*T`.
*   If `S` contains an anonymous field `*T`, the method sets of `S` and `*S` both include promoted methods with receiver `T` or `*T`.

A field declaration may be followed by an optional string literal _tag_, which becomes an attribute for all the fields in the corresponding field declaration. An empty tag string is equivalent to an absent tag. The tags are made visible through a [reflection interface](/pkg/reflect/#StructTag) and take part in [type identity](#Type_identity) for structs but are otherwise ignored.

``` go
struct {
	x, y float64 ""  // an empty tag string is like an absent tag
	name string  "any string is permitted as a tag"
	_    [4]byte "ceci n'est pas un champ de structure"
}

// A struct corresponding to a TimeStamp protocol buffer.
// The tag strings define the protocol buffer field numbers;
// they follow the convention outlined by the reflect package.
struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}
```

### Pointer types ### {#Pointer_types}

A pointer type denotes the set of all pointers to [variables](#Variables) of a given type, called the _base type_ of the pointer. The value of an uninitialized pointer is `nil`.

<pre class="ebnf"><a id="PointerType">PointerType</a> = "*" <a href="#BaseType" class="noline">BaseType</a> .
<a id="BaseType">BaseType</a>    = <a href="#Type" class="noline">Type</a> .
</pre>

```
*Point
*[4]int
```

### Function types ### {#Function_types}

A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is `nil`.

<pre class="ebnf"><a id="FunctionType">FunctionType</a>   = "func" <a href="#Signature" class="noline">Signature</a> .
<a id="Signature">Signature</a>      = <a href="#Parameters" class="noline">Parameters</a> [ <a href="#Result" class="noline">Result</a> ] .
<a id="Result">Result</a>         = <a href="#Parameters" class="noline">Parameters</a> | <a href="#Type" class="noline">Type</a> .
<a id="Parameters">Parameters</a>     = "(" [ <a href="#ParameterList" class="noline">ParameterList</a> [ "," ] ] ")" .
<a id="ParameterList">ParameterList</a>  = <a href="#ParameterDecl" class="noline">ParameterDecl</a> { "," <a href="#ParameterDecl" class="noline">ParameterDecl</a> } .
<a id="ParameterDecl">ParameterDecl</a>  = [ <a href="#IdentifierList" class="noline">IdentifierList</a> ] [ "..." ] <a href="#Type" class="noline">Type</a> .
</pre>

Within a list of parameters or results, the names (IdentifierList) must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-[blank](#Blank_identifier) names in the signature must be [unique](#Uniqueness_of_identifiers). If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

The final incoming parameter in a function signature may have a type prefixed with `...`. A function with such a parameter is called _variadic_ and may be invoked with zero or more arguments for that parameter.

``` go
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```
 
### Interface types ### {#Interface_types}

An interface type specifies a [method set](#Method_sets) called its _interface_. A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to _implement the interface_. The value of an uninitialized variable of interface type is `nil`.

<pre class="ebnf"><a id="InterfaceType">InterfaceType</a>      = "interface" "{" { <a href="#MethodSpec" class="noline">MethodSpec</a> ";" } "}" .
<a id="MethodSpec">MethodSpec</a>         = <a href="#MethodName" class="noline">MethodName</a> <a href="#Signature" class="noline">Signature</a> | <a href="#InterfaceTypeName" class="noline">InterfaceTypeName</a> .
<a id="MethodName">MethodName</a>         = <a href="#identifier" class="noline">identifier</a> .
<a id="InterfaceTypeName">InterfaceTypeName</a>  = <a href="#TypeName" class="noline">TypeName</a> .
</pre>

As with all method sets, in an interface type, each method must have a [unique](#Uniqueness_of_identifiers) non-[blank](#Blank_identifier) name.

``` go
// A simple File interface
interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
	Close()
}
```

More than one type may implement an interface. For instance, if two types `S1` and `S2` have the method set

``` go
func (p T) Read(b Buffer) bool { return … }
func (p T) Write(b Buffer) bool { return … }
func (p T) Close() { … }
```

(where `T` stands for either `S1` or `S2`) then the `File` interface is implemented by both `S1` and `S2`, regardless of what other methods `S1` and `S2` may have or share.

A type implements any interface comprising any subset of its methods and may therefore implement several distinct interfaces. For instance, all types implement the _empty interface_:

``` go
interface{}
```

Similarly, consider this interface specification, which appears within a [type declaration](#Type_declarations) to define an interface called `Locker`:

``` go
type Locker interface {
	Lock()
	Unlock()
}
```

If `S1` and `S2` also implement

``` go
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

they implement the `Locker` interface as well as the `File` interface.

An interface `T` may use a (possibly qualified) interface type name `E` in place of a method specification. This is called _embedding_ interface `E` in `T`; it adds all (exported and non-exported) methods of `E` to the interface `T`.

``` go
type ReadWriter interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type File interface {
	ReadWriter  // same as adding the methods of ReadWriter
	Locker      // same as adding the methods of Locker
	Close()
}

type LockedFile interface {
	Locker
	File        // illegal: Lock, Unlock not unique
	Lock()      // illegal: Lock not unique
}
```

An interface type `T` may not embed itself or any interface type that embeds `T`, recursively.

``` go
// illegal: Bad cannot embed itself
type Bad interface {
	Bad
}

// illegal: Bad1 cannot embed itself using Bad2
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}
```

### Map types ### {#Map_types}

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique _keys_ of another type, called the key type. The value of an uninitialized map is `nil`.

<pre class="ebnf"><a id="MapType">MapType</a>     = "map" "[" <a href="#KeyType" class="noline">KeyType</a> "]" <a href="#ElementType" class="noline">ElementType</a> .
<a id="KeyType">KeyType</a>     = <a href="#Type" class="noline">Type</a> .
</pre>

The [comparison operators](#Comparison_operators) `==` and `!=` must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a [run-time panic](#Run_time_panics).

``` go
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

The number of map elements is called its length. For a map `m`, it can be discovered using the built-in function [`len`](#Length_and_capacity) and may change during execution. Elements may be added during execution using [assignments](#Assignments) and retrieved with [index expressions](#Index_expressions); they may be removed with the [`delete`](#Deletion_of_map_elements) built-in function.

A new, empty map value is made using the built-in function [`make`](#Making_slices_maps_and_channels), which takes the map type and an optional capacity hint as arguments:

``` go
make(map[string]int)
make(map[string]int, 100)
```

The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of `nil` maps. A `nil` map is equivalent to an empty map except that no elements may be added.

### Channel types ### {#Channel_types}

A channel provides a mechanism for [concurrently executing functions](#Go_statements) to communicate by [sending](#Send_statements) and [receiving](#Receive_operator) values of a specified element type. The value of an uninitialized channel is `nil`.

<pre class="ebnf"><a id="ChannelType">ChannelType</a> = ( "chan" | "chan" "&lt;-" | "&lt;-" "chan" ) <a href="#ElementType" class="noline">ElementType</a> .
</pre>

The optional `<-` operator specifies the channel _direction_, _send_ or _receive_. If no direction is given, the channel is _bidirectional_. A channel may be constrained only to send or only to receive by [conversion](#Conversions) or [assignment](#Assignments).

``` go
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```

The `<-` operator associates with the leftmost `chan` possible:

``` go
chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)
```

A new, initialized channel value can be made using the built-in function [`make`](#Making_slices_maps_and_channels), which takes the channel type and an optional _capacity_ as arguments:

``` go
make(chan int, 100)
```

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A `nil` channel is never ready for communication.

A channel may be closed with the built-in function [`close`](#Close). The multi-valued assignment form of the [receive operator](#Receive_operator) reports whether a received value was sent before the channel was closed.

A single channel may be used in [send statements](#Send_statements), [receive operations](#Receive_operator), and calls to the built-in functions [`cap`](#Length_and_capacity) and [`len`](#Length_and_capacity) by any number of goroutines without further synchronization. Channels act as first-in-first-out queues. For example, if one goroutine sends values on a channel and a second goroutine receives them, the values are received in the order sent.
