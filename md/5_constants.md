## Constants ## {#Constants}

There are _boolean constants_, _rune constants_, _integer constants_, _floating-point constants_, _complex constants_, and _string constants_. Rune, integer, floating-point, and complex constants are collectively called _numeric constants_.

A constant value is represented by a [rune](#Rune_literals), [integer](#Integer_literals), [floating-point](#Floating-point_literals), [imaginary](#Imaginary_literals), or [string](#String_literals) literal, an identifier denoting a constant, a [constant expression](#Constant_expressions), a [conversion](#Conversions) with a result that is a constant, or the result value of some built-in functions such as `unsafe.Sizeof` applied to any value, `cap` or `len` applied to [some expressions](#Length_and_capacity), `real` and `imag` applied to a complex constant and `complex` applied to numeric constants. The boolean truth values are represented by the predeclared constants `true` and `false`. The predeclared identifier [iota](#Iota) denotes an integer constant.

In general, complex constants are a form of [constant expression](#Constant_expressions) and are discussed in that section.

Numeric constants represent exact values of arbitrary precision and do not overflow. Consequently, there are no constants denoting the IEEE-754 negative zero, infinity, and not-a-number values.

Constants may be [typed](#Types) or _untyped_. Literal constants, `true`, `false`, `iota`, and certain [constant expressions](#Constant_expressions) containing only untyped constant operands are untyped.

A constant may be given a type explicitly by a [constant declaration](#Constant_declarations) or [conversion](#Conversions), or implicitly when used in a [variable declaration](#Variable_declarations) or an [assignment](#Assignments) or as an operand in an [expression](#Expressions). It is an error if the constant value cannot be represented as a value of the respective type. For instance, `3.0` can be given any integer or any floating-point type, while `2147483648.0` (equal to `1<<31`) can be given the types `float32`, `float64`, or `uint32` but not `int32` or `string`.

An untyped constant has a _default type_ which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a [short variable declaration](#Short_variable_declarations) such as `i := 0` where there is no explicit type. The default type of an untyped constant is `bool`, `rune`, `int`, `float64`, `complex128` or `string` respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.

Implementation restriction: Although numeric constants have arbitrary precision in the language, a compiler may implement them using an internal representation with limited precision. That said, every implementation must:

*   Represent integer constants with at least 256 bits.
*   Represent floating-point constants, including the parts of a complex constant, with a mantissa of at least 256 bits and a signed binary exponent of at least 16 bits.
*   Give an error if unable to represent an integer constant precisely.
*   Give an error if unable to represent a floating-point or complex constant due to overflow.
*   Round to the nearest representable constant if unable to represent a floating-point or complex constant due to limits on precision.

These requirements apply both to literal constants and to the result of evaluating [constant expressions](#Constant_expressions).
