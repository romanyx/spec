## Statements ## {#Statements}

Statements control execution.

<pre class="ebnf"><a id="Statement">Statement</a> =
	<a href="#Declaration" class="noline">Declaration</a> | <a href="#LabeledStmt" class="noline">LabeledStmt</a> | <a href="#SimpleStmt" class="noline">SimpleStmt</a> |
	<a href="#GoStmt" class="noline">GoStmt</a> | <a href="#ReturnStmt" class="noline">ReturnStmt</a> | <a href="#BreakStmt" class="noline">BreakStmt</a> | <a href="#ContinueStmt" class="noline">ContinueStmt</a> | <a href="#GotoStmt" class="noline">GotoStmt</a> |
	<a href="#FallthroughStmt" class="noline">FallthroughStmt</a> | <a href="#Block" class="noline">Block</a> | <a href="#IfStmt" class="noline">IfStmt</a> | <a href="#SwitchStmt" class="noline">SwitchStmt</a> | <a href="#SelectStmt" class="noline">SelectStmt</a> | <a href="#ForStmt" class="noline">ForStmt</a> |
	<a href="#DeferStmt" class="noline">DeferStmt</a> .

<a id="SimpleStmt">SimpleStmt</a> = <a href="#EmptyStmt" class="noline">EmptyStmt</a> | <a href="#ExpressionStmt" class="noline">ExpressionStmt</a> | <a href="#SendStmt" class="noline">SendStmt</a> | <a href="#IncDecStmt" class="noline">IncDecStmt</a> | <a href="#Assignment" class="noline">Assignment</a> | <a href="#ShortVarDecl" class="noline">ShortVarDecl</a> .
</pre>

### Terminating statements ### {#Terminating_statements}

A terminating statement is one of the following:

1.  A ["return"](#Return_statements) or ["goto"](#Goto_statements) statement.
2.  A call to the built-in function [`panic`](#Handling_panics).
3.  A [block](#Blocks) in which the statement list ends in a terminating statement.
4.  An ["if" statement](#If_statements) in which:
    *   the "else" branch is present, and
    *   both branches are terminating statements.
5.  A ["for" statement](#For_statements) in which:
    *   there are no "break" statements referring to the "for" statement, and
    *   the loop condition is absent.
6.  A ["switch" statement](#Switch_statements) in which:
    *   there are no "break" statements referring to the "switch" statement,
    *   there is a default case, and
    *   the statement lists in each case, including the default, end in a terminating statement, or a possibly labeled ["fallthrough" statement](#Fallthrough_statements).
7.  A ["select" statement](#Select_statements) in which:
    *   there are no "break" statements referring to the "select" statement, and
    *   the statement lists in each case, including the default if present, end in a terminating statement.
8.  A [labeled statement](#Labeled_statements) labeling a terminating statement.

All other statements are not terminating.

A [statement list](#Blocks) ends in a terminating statement if the list is not empty and its final non-empty statement is terminating.

### Empty statements ### {#Empty_statements}

The empty statement does nothing.

<pre class="ebnf"><a id="EmptyStmt">EmptyStmt</a> = .
</pre>

### Labeled statements ### {#Labeled_statements}

A labeled statement may be the target of a `goto`, `break` or `continue` statement.

<pre class="ebnf"><a id="LabeledStmt">LabeledStmt</a> = <a href="#Label" class="noline">Label</a> ":" <a href="#Statement" class="noline">Statement</a> .
<a id="Label">Label</a>       = <a href="#identifier" class="noline">identifier</a> .
</pre>

``` go
Error: log.Panic("error encountered")
```

### Expression statements ### {#Expression_statements}

With the exception of specific built-in functions, function and method [calls](#Calls) and [receive operations](#Receive_operator) can appear in statement context. Such statements may be parenthesized.

<pre class="ebnf"><a id="ExpressionStmt">ExpressionStmt</a> = <a href="#Expression" class="noline">Expression</a> .
</pre>

The following built-in functions are not permitted in statement context:

``` go
append cap complex imag len make new real
unsafe.Alignof unsafe.Offsetof unsafe.Sizeof
```

``` go
h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function
```

### Send statements ### {#Send_statements}

A send statement sends a value on a channel. The channel expression must be of [channel type](#Channel_types), the channel direction must permit send operations, and the type of the value to be sent must be [assignable](#Assignability) to the channel's element type.

<pre class="ebnf"><a id="SendStmt">SendStmt</a> = <a href="#Channel" class="noline">Channel</a> "&lt;-" <a href="#Expression" class="noline">Expression</a> .
<a id="Channel">Channel</a>  = <a href="#Expression" class="noline">Expression</a> .
</pre>

Both the channel and the value expression are evaluated before communication begins. Communication blocks until the send can proceed. A send on an unbuffered channel can proceed if a receiver is ready. A send on a buffered channel can proceed if there is room in the buffer. A send on a closed channel proceeds by causing a [run-time panic](#Run_time_panics). A send on a `nil` channel blocks forever.

``` go
ch <- 3  // send value 3 to channel ch
```

### IncDec statements ### {#IncDec_statements}

The "++" and "--" statements increment or decrement their operands by the untyped [constant](#Constants) `1`. As with an assignment, the operand must be [addressable](#Address_operators) or a map index expression.

<pre class="ebnf"><a id="IncDecStmt">IncDecStmt</a> = <a href="#Expression" class="noline">Expression</a> ( "++" | "--" ) .
</pre>

The following [assignment statements](#Assignments) are semantically equivalent:

``` go
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
```

### Assignments ### {#Assignments}

<pre class="ebnf"><a id="Assignment">Assignment</a> = <a href="#ExpressionList" class="noline">ExpressionList</a> <a href="#assign_op" class="noline">assign_op</a> <a href="#ExpressionList" class="noline">ExpressionList</a> .

<a id="assign_op">assign_op</a> = [ <a href="#add_op" class="noline">add_op</a> | <a href="#mul_op" class="noline">mul_op</a> ] "=" .
</pre>

Each left-hand side operand must be [addressable](#Address_operators), a map index expression, or (for `=` assignments only) the [blank identifier](#Blank_identifier). Operands may be parenthesized.

``` go
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch
```

An _assignment operation_ `x` _op_`=` `y` where _op_ is a binary arithmetic operation is equivalent to `x` `=` `x` _op_ `(y)` but evaluates `x` only once. The _op_`=` construct is a single token. In assignment operations, both the left- and right-hand expression lists must contain exactly one single-valued expression, and the left-hand expression must not be the blank identifier.

``` go
a[i] <<= 2
i &^= 1<<n
```

A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. There are two forms. In the first, the right hand operand is a single multi-valued expression such as a function call, a [channel](#Channel_types) or [map](#Map_types) operation, or a [type assertion](#Type_assertions). The number of operands on the left hand side must match the number of values. For instance, if `f` is a function returning two values,

``` go
x, y = f()
```

assigns the first value to `x` and the second to `y`. In the second form, the number of operands on the left must equal the number of expressions on the right, each of which must be single-valued, and the _n_th expression on the right is assigned to the _n_th operand on the left:

``` go
one, two, three = '一', '二', '三'
```

The [blank identifier](#Blank_identifier) provides a way to ignore right-hand side values in an assignment:

``` go
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
```

The assignment proceeds in two phases. First, the operands of [index expressions](#Index_expressions) and [pointer indirections](#Address_operators) (including implicit pointer indirections in [selectors](#Selectors)) on the left and the expressions on the right are all [evaluated in the usual order](#Order_of_evaluation). Second, the assignments are carried out in left-to-right order.

``` go
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x == []int{3, 5, 3}
```

In assignments, each value must be [assignable](#Assignability) to the type of the operand to which it is assigned, with the following special cases:

1.  Any typed value may be assigned to the blank identifier.
2.  If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first [converted](#Conversions) to its [default type](#Constants).
3.  If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first converted to type `bool`.

### If statements ### {#If_statements}

"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.

<pre class="ebnf"><a id="IfStmt">IfStmt</a> = "if" [ <a href="#SimpleStmt" class="noline">SimpleStmt</a> ";" ] <a href="#Expression" class="noline">Expression</a> <a href="#Block" class="noline">Block</a> [ "else" ( <a href="#IfStmt" class="noline">IfStmt</a> | <a href="#Block" class="noline">Block</a> ) ] .
</pre>

``` go
if x > max {
	x = max
}
```

The expression may be preceded by a simple statement, which executes before the expression is evaluated.

``` go
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

### Switch statements ### {#Switch_statements}

"Switch" statements provide multi-way execution. An expression or type specifier is compared to the "cases" inside the "switch" to determine which branch to execute.

<pre class="ebnf"><a id="SwitchStmt">SwitchStmt</a> = <a href="#ExprSwitchStmt" class="noline">ExprSwitchStmt</a> | <a href="#TypeSwitchStmt" class="noline">TypeSwitchStmt</a> .
</pre>

There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression. In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. The switch expression is evaluated exactly once in a switch statement.

#### Expression switches

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value `true`.

<pre class="ebnf"><a id="ExprSwitchStmt">ExprSwitchStmt</a> = "switch" [ <a href="#SimpleStmt" class="noline">SimpleStmt</a> ";" ] [ <a href="#Expression" class="noline">Expression</a> ] "{" { <a href="#ExprCaseClause" class="noline">ExprCaseClause</a> } "}" .
<a id="ExprCaseClause">ExprCaseClause</a> = <a href="#ExprSwitchCase" class="noline">ExprSwitchCase</a> ":" <a href="#StatementList" class="noline">StatementList</a> .
<a id="ExprSwitchCase">ExprSwitchCase</a> = "case" <a href="#ExpressionList" class="noline">ExpressionList</a> | "default" .
</pre>

If the switch expression evaluates to an untyped constant, it is first [converted](#Conversions) to its [default type](#Constants); if it is an untyped boolean value, it is first converted to type `bool`. The predeclared untyped value `nil` cannot be used as a switch expression.

If a case expression is untyped, it is first [converted](#Conversions) to the type of the switch expression. For each (possibly converted) case expression `x` and the value `t` of the switch expression, `x == t` must be a valid [comparison](#Comparison_operators).

In other words, the switch expression is treated as if it were used to declare and initialize a temporary variable `t` without explicit type; it is that value of `t` against which each case expression `x` is tested for equality.

In a case or default clause, the last non-empty statement may be a (possibly [labeled](#Labeled_statements)) ["fallthrough" statement](#Fallthrough_statements) to indicate that control should flow from the end of this clause to the first statement of the next clause. Otherwise control flows to the end of the "switch" statement. A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.

The switch expression may be preceded by a simple statement, which executes before the expression is evaluated.

``` go
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```

Implementation restriction: A compiler may disallow multiple case expressions evaluating to the same constant. For instance, the current compilers disallow duplicate integer, floating point, or string constants in case expressions.

#### Type switches

A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a [type assertion](#Type_assertions) using the reserved word `type` rather than an actual type:

``` go
switch x.(type) {
// cases
}
```

Cases then match actual types `T` against the dynamic type of the expression `x`. As with type assertions, `x` must be of [interface type](#Interface_types), and each non-interface type `T` listed in a case must implement the type of `x`. The types listed in the cases of a type switch must all be [different](#Type_identity).

<pre class="ebnf"><a id="TypeSwitchStmt">TypeSwitchStmt</a>  = "switch" [ <a href="#SimpleStmt" class="noline">SimpleStmt</a> ";" ] <a href="#TypeSwitchGuard" class="noline">TypeSwitchGuard</a> "{" { <a href="#TypeCaseClause" class="noline">TypeCaseClause</a> } "}" .
<a id="TypeSwitchGuard">TypeSwitchGuard</a> = [ <a href="#identifier" class="noline">identifier</a> ":=" ] <a href="#PrimaryExpr" class="noline">PrimaryExpr</a> "." "(" "type" ")" .
<a id="TypeCaseClause">TypeCaseClause</a>  = <a href="#TypeSwitchCase" class="noline">TypeSwitchCase</a> ":" <a href="#StatementList" class="noline">StatementList</a> .
<a id="TypeSwitchCase">TypeSwitchCase</a>  = "case" <a href="#TypeList" class="noline">TypeList</a> | "default" .
<a id="TypeList">TypeList</a>        = <a href="#Type" class="noline">Type</a> { "," <a href="#Type" class="noline">Type</a> } .
</pre>

The TypeSwitchGuard may include a [short variable declaration](#Short_variable_declarations). When that form is used, the variable is declared at the end of the TypeSwitchCase in the [implicit block](#Blocks) of each clause. In clauses with a case listing exactly one type, the variable has that type; otherwise, the variable has the type of the expression in the TypeSwitchGuard.

The type in a case may be [`nil`](#Predeclared_identifiers); that case is used when the expression in the TypeSwitchGuard is a `nil` interface value. There may be at most one `nil` case.

Given an expression `x` of type `interface{}`, the following type switch:

``` go
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```

could be rewritten:

``` go
v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
```

The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.

The "fallthrough" statement is not permitted in a type switch.

### For statements ### {#For_statements}

A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.

<pre class="ebnf"><a id="ForStmt">ForStmt</a> = "for" [ <a href="#Condition" class="noline">Condition</a> | <a href="#ForClause" class="noline">ForClause</a> | <a href="#RangeClause" class="noline">RangeClause</a> ] <a href="#Block" class="noline">Block</a> .
<a id="Condition">Condition</a> = <a href="#Expression" class="noline">Expression</a> .
</pre>

#### For statements with single condition

In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value `true`.

``` go
for a < b {
	a *= 2
}
```

#### For statements with `for` clause

A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an _init_ and a _post_ statement, such as an assignment, an increment or decrement statement. The init statement may be a [short variable declaration](#Short_variable_declarations), but the post statement must not. Variables declared by the init statement are re-used in each iteration.

<pre class="ebnf"><a id="ForClause">ForClause</a> = [ <a href="#InitStmt" class="noline">InitStmt</a> ] ";" [ <a href="#Condition" class="noline">Condition</a> ] ";" [ <a href="#PostStmt" class="noline">PostStmt</a> ] .
<a id="InitStmt">InitStmt</a> = <a href="#SimpleStmt" class="noline">SimpleStmt</a> .
<a id="PostStmt">PostStmt</a> = <a href="#SimpleStmt" class="noline">SimpleStmt</a> .
</pre>
``` go
for i := 0; i < 10; i++ {
	f(i)
}
```

If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the [semicolons](#Semicolons) are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value `true`.

``` go
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
```

#### For statements with `range` clause

A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel. For each entry it assigns _iteration values_ to corresponding _iteration variables_ if present and then executes the block.

<pre class="ebnf"><a id="RangeClause">RangeClause</a> = [ <a href="#ExpressionList" class="noline">ExpressionList</a> "=" | <a href="#IdentifierList" class="noline">IdentifierList</a> ":=" ] "range" <a href="#Expression" class="noline">Expression</a> .
</pre>

The expression on the right in the "range" clause is called the _range expression_, which may be an array, pointer to an array, slice, string, map, or channel permitting [receive operations](#Receive_operator). As with an assignment, if present the operands on the left must be [addressable](#Address_operators) or map index expressions; they denote the iteration variables. If the range expression is a channel, at most one iteration variable is permitted, otherwise there may be up to two. If the last iteration variable is the [blank identifier](#Blank_identifier), the range clause is equivalent to the same clause without that identifier.

The range expression is evaluated once before beginning the loop, with one exception: if the range expression is an array or a pointer to an array and at most one iteration variable is present, only the range expression's length is evaluated; if that length is constant, [by definition](#Length_and_capacity) the range expression itself will not be evaluated.

Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present:

``` go
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```

1.  For an array, pointer to array, or slice value `a`, the index iteration values are produced in increasing order, starting at element index 0. If at most one iteration variable is present, the range loop produces iteration values from 0 up to `len(a)-1` and does not index into the array or slice itself. For a `nil` slice, the number of iterations is 0.
2.  For a string value, the "range" clause iterates over the Unicode code points in the string starting at byte index 0\. On successive iterations, the index value will be the index of the first byte of successive UTF-8-encoded code points in the string, and the second value, of type `rune`, will be the value of the corresponding code point. If the iteration encounters an invalid UTF-8 sequence, the second value will be `0xFFFD`, the Unicode replacement character, and the next iteration will advance a single byte in the string.
3.  The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next. If map entries that have not yet been reached are removed during iteration, the corresponding iteration values will not be produced. If map entries are created during iteration, that entry may be produced during the iteration or may be skipped. The choice may vary for each entry created and from one iteration to the next. If the map is `nil`, the number of iterations is 0.
4.  For channels, the iteration values produced are the successive values sent on the channel until the channel is [closed](#Close). If the channel is `nil`, the range expression blocks forever.

The iteration values are assigned to the respective iteration variables as in an [assignment statement](#Assignments).

The iteration variables may be declared by the "range" clause using a form of [short variable declaration](#Short_variable_declarations) (`:=`). In this case their types are set to the types of the respective iteration values and their [scope](#Declarations_and_scope) is the block of the "for" statement; they are re-used in each iteration. If the iteration variables are declared outside the "for" statement, after execution their values will be those of the last iteration.

``` go
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface {}  // value type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel
for range ch {}
```

### Go statements ### {#Go_statements}

A "go" statement starts the execution of a function call as an independent concurrent thread of control, or _goroutine_, within the same address space.

<pre class="ebnf"><a id="GoStmt">GoStmt</a> = "go" <a href="#Expression" class="noline">Expression</a> .
</pre>

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](#Expression_statements).

The function value and parameters are [evaluated as usual](#Calls) in the calling goroutine, but unlike with a regular call, program execution does not wait for the invoked function to complete. Instead, the function begins executing independently in a new goroutine. When the function terminates, its goroutine also terminates. If the function has any return values, they are discarded when the function completes.

``` go
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true; }} (c)
```

### Select statements ### {#Select_statements}

A "select" statement chooses which of a set of possible [send](#Send_statements) or [receive](#Receive_operator) operations will proceed. It looks similar to a ["switch"](#Switch_statements) statement but with the cases all referring to communication operations.

<pre class="ebnf"><a id="SelectStmt">SelectStmt</a> = "select" "{" { <a href="#CommClause" class="noline">CommClause</a> } "}" .
<a id="CommClause">CommClause</a> = <a href="#CommCase" class="noline">CommCase</a> ":" <a href="#StatementList" class="noline">StatementList</a> .
<a id="CommCase">CommCase</a>   = "case" ( <a href="#SendStmt" class="noline">SendStmt</a> | <a href="#RecvStmt" class="noline">RecvStmt</a> ) | "default" .
<a id="RecvStmt">RecvStmt</a>   = [ <a href="#ExpressionList" class="noline">ExpressionList</a> "=" | <a href="#IdentifierList" class="noline">IdentifierList</a> ":=" ] <a href="#RecvExpr" class="noline">RecvExpr</a> .
<a id="RecvExpr">RecvExpr</a>   = <a href="#Expression" class="noline">Expression</a> .
</pre>

A case with a RecvStmt may assign the result of a RecvExpr to one or two variables, which may be declared using a [short variable declaration](#Short_variable_declarations). The RecvExpr must be a (possibly parenthesized) receive operation. There can be at most one default case and it may appear anywhere in the list of cases.

Execution of a "select" statement proceeds in several steps:

1.  For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.
2.  If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
3.  Unless the selected case is the default case, the respective communication operation is executed.
4.  If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned.
5.  The statement list of the selected case is executed.

Since communication on `nil` channels can never proceed, a select with only `nil` channels and no default case blocks forever.

``` go
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = <-c1:
	print("received ", i1, " from c1\n")
case c2 <- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
case a[f()] = <-c4:
	// same as:
	// case t := <-c4
	//	a[f()] = t
default:
	print("no communication\n")
}

for {  // send random sequence of bits to c
	select {
	case c <- 0:  // note: no statement, no fallthrough, no folding of cases
	case c <- 1:
	}
}

select {}  // block forever
```

### Return statements ### {#Return_statements}

A "return" statement in a function `F` terminates the execution of `F`, and optionally provides one or more result values. Any functions [deferred](#Defer_statements) by `F` are executed before `F` returns to its caller.

<pre class="ebnf"><a id="ReturnStmt">ReturnStmt</a> = "return" [ <a href="#ExpressionList" class="noline">ExpressionList</a> ] .
</pre>

In a function without a result type, a "return" statement must not specify any result values.

``` go
func noResult() {
	return
}
```

There are three ways to return values from a function with a result type:

1.  The return value or values may be explicitly listed in the "return" statement. Each expression must be single-valued and [assignable](#Assignability) to the corresponding element of the function's result type.

    <pre>func simpleF() int {
    	return 2
    }

    func complexF1() (re float64, im float64) {
    	return -7.0, -4.0
    }
    </pre>

2.  The expression list in the "return" statement may be a single call to a multi-valued function. The effect is as if each value returned from that function were assigned to a temporary variable with the type of the respective value, followed by a "return" statement listing these variables, at which point the rules of the previous case apply.

    <pre>func complexF2() (re float64, im float64) {
    	return complexF1()
    }
	</pre>

3.  The expression list may be empty if the function's result type specifies names for its [result parameters](#Function_types). The result parameters act as ordinary local variables and the function may assign values to them as necessary. The "return" statement returns the values of these variables.

	<pre>func complexF3() (re float64, im float64) {
			re = 7.0
			im = 4.0
			return
		}

		func (devnull) Write(p []byte) (n int, _ error) {
			n = len(p)
			return
	</pre>

Regardless of how they are declared, all the result values are initialized to the [zero values](#The_zero_value) for their type upon entry to the function. A "return" statement that specifies results sets the result parameters before any deferred functions are executed.

Implementation restriction: A compiler may disallow an empty expression list in a "return" statement if a different entity (constant, type, or variable) with the same name as a result parameter is in [scope](#Declarations_and_scope) at the place of the return.

``` go
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed
	}
	return
}
```

### Break statements ### {#Break_statements}

A "break" statement terminates execution of the innermost ["for"](#For_statements), ["switch"](#Switch_statements), or ["select"](#Select_statements) statement within the same function.

<pre class="ebnf"><a id="BreakStmt">BreakStmt</a> = "break" [ <a href="#Label" class="noline">Label</a> ] .
</pre>

If there is a label, it must be that of an enclosing "for", "switch", or "select" statement, and that is the one whose execution terminates.

``` go
OuterLoop:
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
```

### Continue statements ### {#Continue_statements}

A "continue" statement begins the next iteration of the innermost ["for" loop](#For_statements) at its post statement. The "for" loop must be within the same function.

<pre class="ebnf"><a id="ContinueStmt">ContinueStmt</a> = "continue" [ <a href="#Label" class="noline">Label</a> ] .
</pre>

If there is a label, it must be that of an enclosing "for" statement, and that is the one whose execution advances.

``` go
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
```

### Goto statements ### {#Goto_statements}

A "goto" statement transfers control to the statement with the corresponding label within the same function.

<pre class="ebnf"><a id="GotoStmt">GotoStmt</a> = "goto" <a href="#Label" class="noline">Label</a> .
</pre>

``` go
goto Error
```

Executing the "goto" statement must not cause any variables to come into [scope](#Declarations_and_scope) that were not already in scope at the point of the goto. For instance, this example:

``` go
	goto L  // BAD
	v := 3
L:
```

is erroneous because the jump to label `L` skips the creation of `v`.

A "goto" statement outside a [block](#Blocks) cannot jump to a label inside that block. For instance, this example:

``` go
if n%2 == 1 {
	goto L1
}
for n > 0 {
	f()
	n--
L1:
	f()
	n--
}
```

is erroneous because the label `L1` is inside the "for" statement's block but the `goto` is not.

### Fallthrough statements ### {#Fallthrough_statements}

A "fallthrough" statement transfers control to the first statement of the next case clause in an [expression "switch" statement](#Expression_switches). It may be used only as the final non-empty statement in such a clause.

<pre class="ebnf"><a id="FallthroughStmt">FallthroughStmt</a> = "fallthrough" .
</pre>

### Defer statements ### {#Defer_statements}

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a [return statement](#Return_statements), reached the end of its [function body](#Function_declarations), or because the corresponding goroutine is [panicking](#Handling_panics).

<pre class="ebnf"><a id="DeferStmt">DeferStmt</a> = "defer" <a href="#Expression" class="noline">Expression</a> .
</pre>

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](#Expression_statements).

Each time a "defer" statement executes, the function value and parameters to the call are [evaluated as usual](#Calls) and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. If a deferred function value evaluates to `nil`, execution [panics](#Handling_panics) when the function is invoked, not when the "defer" statement is executed.

For instance, if the deferred function is a [function literal](#Function_literals) and the surrounding function has [named result parameters](#Function_types) that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes. (See also the section on [handling panics](#Handling_panics).)

``` go
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// prints 3 2 1 0 before surrounding function returns
for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

// f returns 1
func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}
```
