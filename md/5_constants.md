## Константы ## {#Constants}

В языке присутствуют: _булевы константы_, _рунные константы_, _целочисленные константы_, _константы чисел с плавающей точкой_, _константы комплексных чисел_ и _строковые константы_. Рунные, целочисленные и константы с плавающей точкой вместе называются _числовыми константами_.

Значение константы представлено [рунным](#Rune_literals), [целочисленным](#Integer_literals), [мнимым](#Imaginary_literals), [строчным](#String_literals) литералом или литералом [числа с плавающей точкой](#Floating-point_literals), идентификатором обозначающим константу, [выражением-константой](#Constant_expressions), [преобразованием](#Conversions), результатом которого будет константа, или результирующим значением некоторых встроенных функций, таких как `unsafe.Sizeof` применимой к любому значению, `cap` или `len` применимым к [некоторым выражениям](#Length_and_capacity), `real` и `imag` применимым к константе комплексного числа, и `complex` применимой к числовым константам. Булевы значения истинности представлены предопределенными константами `true` и `false`. Предопределенный идентификатор [iota](#Iota) обозначаем целочисленную константу.

В общем константы комплексных чисел являются одной из форм [выражений-констант](#Constant_expressions) и будут детально рассмотрены в соответствующем разделе.

Численные константы представляют собой точные значения произвольной разрядности, которые не переполняются. Следовательно, отсутствуют константы представляющие значения отрицательного нуля, бесконечности и _нечисла_ из стандарта [IEEE-754](https://ru.wikipedia.org/wiki/IEEE_754-2008).

Константы могут быть [типизированными](#Types) и _нетипизированными_. Литеральные константы, `true`, `false`, `iota` и некоторые [выражения-константы](#Constant_expressions), содержащие только нетипизированные операнды-константы, являются нетипизированными.

A constant may be given a type explicitly by a [constant declaration](#Constant_declarations) or [conversion](#Conversions), or implicitly when used in a [variable declaration](#Variable_declarations) or an [assignment](#Assignments) or as an operand in an [expression](#Expressions). It is an error if the constant value cannot be represented as a value of the respective type. For instance, `3.0` can be given any integer or any floating-point type, while `2147483648.0` (equal to `1<<31`) can be given the types `float32`, `float64`, or `uint32` but not `int32` or `string`.

An untyped constant has a _default type_ which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a [short variable declaration](#Short_variable_declarations) such as `i := 0` where there is no explicit type. The default type of an untyped constant is `bool`, `rune`, `int`, `float64`, `complex128` or `string` respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.

Implementation restriction: Although numeric constants have arbitrary precision in the language, a compiler may implement them using an internal representation with limited precision. That said, every implementation must:

*   Represent integer constants with at least 256 bits.
*   Represent floating-point constants, including the parts of a complex constant, with a mantissa of at least 256 bits and a signed binary exponent of at least 16 bits.
*   Give an error if unable to represent an integer constant precisely.
*   Give an error if unable to represent a floating-point or complex constant due to overflow.
*   Round to the nearest representable constant if unable to represent a floating-point or complex constant due to limits on precision.

These requirements apply both to literal constants and to the result of evaluating [constant expressions](#Constant_expressions).
