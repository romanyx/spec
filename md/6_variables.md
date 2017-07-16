## Variables ## {#Variables}

A variable is a storage location for holding a _value_. The set of permissible values is determined by the variable's _[type](#Types)_.

A [variable declaration](#Variable_declarations) or, for function parameters and results, the signature of a [function declaration](#Function_declarations) or [function literal](#Function_literals) reserves storage for a named variable. Calling the built-in function [`new`](#Allocation) or taking the address of a [composite literal](#Composite_literals) allocates storage for a variable at run time. Such an anonymous variable is referred to via a (possibly implicit) [pointer indirection](#Address_operators).

_Structured_ variables of [array](#Array_types), [slice](#Slice_types), and [struct](#Struct_types) types have elements and fields that may be [addressed](#Address_operators) individually. Each such element acts like a variable.

The _static type_ (or just _type_) of a variable is the type given in its declaration, the type provided in the `new` call or composite literal, or the type of an element of a structured variable. Variables of interface type also have a distinct _dynamic type_, which is the concrete type of the value assigned to the variable at run time (unless the value is the predeclared identifier `nil`, which has no type). The dynamic type may vary during execution but values stored in interface variables are always [assignable](#Assignability) to the static type of the variable.

``` go
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
```

A variable's value is retrieved by referring to the variable in an [expression](#Expressions); it is the most recent value [assigned](#Assignments) to the variable. If a variable has not yet been assigned a value, its value is the [zero value](#The_zero_value) for its type.
