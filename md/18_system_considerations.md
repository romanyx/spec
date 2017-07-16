## System considerations ## {#System_considerations}

### Package `unsafe` ### {#Package_unsafe}

The built-in package `unsafe`, known to the compiler, provides facilities for low-level programming including operations that violate the type system. A package using `unsafe` must be vetted manually for type safety and may not be portable. The package provides the following interface:

``` go
package unsafe

type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
type Pointer *ArbitraryType

func Alignof(variable ArbitraryType) uintptr
func Offsetof(selector ArbitraryType) uintptr
func Sizeof(variable ArbitraryType) uintptr
```

A `Pointer` is a [pointer type](#Pointer_types) but a `Pointer` value may not be [dereferenced](#Address_operators). Any pointer or value of [underlying type](#Types) `uintptr` can be converted to a `Pointer` type and vice versa. The effect of converting between `Pointer` and `uintptr` is implementation-defined.

``` go
var f float64
bits = *(*uint64)(unsafe.Pointer(&f))

type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&f))

var p ptr = nil
```

The functions `Alignof` and `Sizeof` take an expression `x` of any type and return the alignment or size, respectively, of a hypothetical variable `v` as if `v` was declared via `var v = x`.

The function `Offsetof` takes a (possibly parenthesized) [selector](#Selectors) `s.f`, denoting a field `f` of the struct denoted by `s` or `*s`, and returns the field offset in bytes relative to the struct's address. If `f` is an [embedded field](#Struct_types), it must be reachable without pointer indirections through fields of the struct. For a struct `s` with field `f`:

``` go
uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&s.f))
```

Computer architectures may require memory addresses to be _aligned_; that is, for addresses of a variable to be a multiple of a factor, the variable's type's _alignment_. The function `Alignof` takes an expression denoting a variable of any type and returns the alignment of the (type of the) variable in bytes. For a variable `x`:

``` go
uintptr(unsafe.Pointer(&x)) % unsafe.Alignof(x) == 0
```

Calls to `Alignof`, `Offsetof`, and `Sizeof` are compile-time constant expressions of type `uintptr`.

### Size and alignment guarantees ### {##Size_and_alignment_guarantees}

For the [numeric types](#Numeric_types), the following sizes are guaranteed:

``` go
type                                 size in bytes

byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16
```

The following minimal alignment properties are guaranteed:

1.  For a variable `x` of any type: `unsafe.Alignof(x)` is at least 1.
2.  For a variable `x` of struct type: `unsafe.Alignof(x)` is the largest of all the values `unsafe.Alignof(x.f)` for each field `f` of `x`, but at least 1.
3.  For a variable `x` of array type: `unsafe.Alignof(x)` is the same as `unsafe.Alignof(x[0])`, but at least 1.

A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. Two distinct zero-size variables may have the same address in memory.
