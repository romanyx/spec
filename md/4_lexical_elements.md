## Лексические элементы ## {#Lexical_elements}

### Комментарии ### {#Comments}

Комментарии используются для документирования кода. Существуют два способа обозначить комментарий:

1. _Строчные комментарии_ начинаются с последовательности символов `//` и ограничиваются строкой
2. _Блочные комментарии_ начинаются с символов `/*` и закрываются первой встреченной последовательностью символов `*/`

Комментарии не могу начинаться внутри [руны](#Rune_literals), [строкового литерала](#String_literals) или внутри комментария. Блочные комментарии расположенные на одной строке соединяются между собой пробелом, все другие разделяются новой строкой.

### Токены ### {#Tokens}

Токены формируют словарь языка Go. Все они разделены на четыре класса: _идентификаторы_, _ключевые слова_, _операторы и разделители_ и _литералы_. _Пустое пространство_, образованное символами пробела (U+0020) табуляция (U+0009), возврат каретки (U+000D) и новая строка (U+000A) игнорируются за исключением случаев когда они разделяют токены, которые в противном случае воспринимались как один. Помимо этого символ новой строки или конец файла могут инициировать вставку [точки с запятой](#Semicolons) лексическим анализатором. При разбиении ввода на токены, последующий токен будет являться самым длинным набором символов, воспринимаемым как действительный токен.

### Точка с запятой ### {#Semicolons}

Формальная грамматика Go использует точку с запятой `";"` для терминирования результатов. Однако исходный код написанный на Go в большенстве случаев может упускать точку с запятой, придерживаясь следующих правил.

1. Когда ввод разбит на токены, точка с запятой автоматически разместиться за последним токеном во входящей строке если токен является

    * [Идентификатором](#Identifiers)
    * Литералом [целочисленного числа](#Integer_literals), литералом [числа с плавающей точкой](#Floating-point_literals), [рунным](#Rune_literals) литералом, [строковым](#String_literals) литералом или [мнимым](#Imaginary_literals) литералом.
    * [Ключевым словом](#Keywords) `break`, `continue`, `fallthrough` или `return`
    * Одним из [операторов или разделителей](#Operators_and_Delimiters) `++`, `--`, `)`, `]`, `}`

2. Для возможности записи сложных конструкций в одну строку, точка с запятой будет вставлена перед закрывающими круглой `")"` и фигурной `"}"` скобками

Чтобы показать идиоматическое использование точки с запятой, примеры кода приведенные в данной документации используют приведенные правила

### Идентификаторы ### {#Identifiers}

Идентификаторы определяют названия программных сущностей таких как переменные и типы. Идентификаторы являются последовательностью букв и цифр. Первым символом идентификатора должна быть буква

<pre class="ebnf"><a id="identifier">identifier</a> = <a href="#letter" class="noline">letter</a> { <a href="#letter" class="noline">letter</a> | <a href="#unicode_digit" class="noline">unicode_digit</a> } .
</pre>

``` go
a
_x9
ThisVariableIsExported
αβ
```

Часть идентификаторов [предопределена](#Predeclared_identifiers) заранее.

### Ключевые слова ### {#Keywords}

Следующие ключевые слова зарезервированы и не могут использоватся как идентификаторы

``` go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Операторы и разделители ### {#Operators_and_Delimiters}

Следующие последовательности символов представляют [операторы](#Operators), разделители и другие специальные токены

``` go
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

### Целочисленные литералы ### {#Integer_literals}

Целочисленные литералы это последовательность цифр представляющая собой [целочисленную константу](#Constants). Для обозначения не десятичных чисел используется префиксы: `0` для восьмеричной, `0x` или `0X` для шестеричной. В шестнадцатеричных литералах, буквы `a-f` и `A-F` обозначают значения в диапазоне 10-15.

<pre class="ebnf"><a id="int_lit">int_lit</a>     = <a href="#decimal_lit" class="noline">decimal_lit</a> | <a href="#octal_lit" class="noline">octal_lit</a> | <a href="#hex_lit" class="noline">hex_lit</a> .
<a id="decimal_lit">decimal_lit</a> = ( "1" … "9" ) { <a href="#decimal_digit" class="noline">decimal_digit</a> } .
<a id="octal_lit">octal_lit</a>   = "0" { <a href="#octal_digit" class="noline">octal_digit</a> } .
<a id="hex_lit">hex_lit</a>     = "0" ( "x" | "X" ) <a href="#hex_digit" class="noline">hex_digit</a> { <a href="#hex_digit" class="noline">hex_digit</a> } .
</pre>

``` go
42
0600
0xBadFace
170141183460469231731687303715884105727
```

### Литералы чисел с плавающей точкой ### {#Floating-point_literals}

Литерал с плавающей точкой это десятеричное представление констант чисел с [плавающей точкой](###). Они имеют целую часть, десятичную точку, дробную часть и экспоненциальную часть. Целая и дробная часть образуют десятичные цифры. Экспонента представлена символом `e` или `E` за которым следует необязательный десятичный показатель. В десятичной записи может быть упущена целая или дробная часть, в экспонентной записи может быть упущена десятичная точка или показатель степени

<pre class="ebnf"><a id="float_lit">float_lit</a> = <a href="#decimals" class="noline">decimals</a> "." [ <a href="#decimals" class="noline">decimals</a> ] [ <a href="#exponent" class="noline">exponent</a> ] |
            <a href="#decimals" class="noline">decimals</a> <a href="#exponent" class="noline">exponent</a> |
            "." <a href="#decimals" class="noline">decimals</a> [ <a href="#exponent" class="noline">exponent</a> ] .
<a id="decimals">decimals</a>  = <a href="#decimal_digit" class="noline">decimal_digit</a> { <a href="#decimal_digit" class="noline">decimal_digit</a> } .
<a id="exponent">exponent</a>  = ( "e" | "E" ) [ "+" | "-" ] <a href="#decimals" class="noline">decimals</a> .
</pre>

``` go
0.
72.40
072.40  // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
```

### Мнимые литералы ### {#Imaginary_literals}

Мнимый литерал это десятеричное представление мнимой части [констант комплексных чисел](#Constants). Он состоит из литерала [числа с плавающей точкой](#Floating-point_literals) или десятичного целого числа с символом `i` в конце

<pre class="ebnf"><a id="imaginary_lit">imaginary_lit</a> = (<a href="#decimals" class="noline">decimals</a> | <a href="#float_lit" class="noline">float_lit</a>) "i" .
</pre>

``` go
0i
011i  // == 11i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
```

### Рунные литералы ### {#Rune_literals}

Рунные литералы представляют [рунные константы](#Constants), целочисленные значения представляющие кодовые точки Unicode. Рунный литерал записывается как один или несколько символов размещенных в одинарные кавычки, как пример `'x'`, `'\n'`. Внутри одинарных кавычек может быть размещен любой символ кроме новой строки и неэкранированной одинарной кавычки. Одиночный символ представляет собой непосредственно кодовую точку Unicode, в то время как мультисимвольная запись может начинающаяся с обратного слэша, кодирует значение в различных форматах.

Простая форма записи представляет собой одиночным символом заключенный в одинарные кавычки. Так как исходный код в Go представляет собой набор кодовых точек Unicode закодированных в UTF-8, несколько байт закодированных в UTF-8 могут представлять один символ. Для примера знак `'a'` содержит единственный байт представляющий букву `a`, кодовую точку Unicode U+0061, шестеричное значение `0x61`, тогда как литерал `'ä'` представляет собой два байта (`0xc3` `0xa4`) представляющих `a` с тремой, кодовую точку Unicode U+00E4 и шестеричное значение `0xe4`

Several backslash escapes allow arbitrary values to be encoded as ASCII text. There are four ways to represent the integer value as a numeric constant: `\x` followed by exactly two hexadecimal digits; `\u` followed by exactly four hexadecimal digits; `\U` followed by exactly eight hexadecimal digits, and a plain backslash `\` followed by exactly three octal digits. In each case the value of the literal is the value represented by the digits in the corresponding base.

Although these representations all result in an integer, they have different valid ranges. Octal escapes must represent a value between 0 and 255 inclusive. Hexadecimal escapes satisfy this condition by construction. The escapes `\u` and `\U` represent Unicode code points so within them some values are illegal, in particular those above `0x10FFFF` and surrogate halves.

After a backslash, certain single-character escapes represent special values:

``` go
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000b vertical tab
\\   U+005c backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
```

All other sequences starting with a backslash are illegal inside rune literals.

<pre class="ebnf"><a id="rune_lit">rune_lit</a>         = "'" ( <a href="#unicode_value" class="noline">unicode_value</a> | <a href="#byte_value" class="noline">byte_value</a> ) "'" .
<a id="unicode_value">unicode_value</a>    = <a href="#unicode_char" class="noline">unicode_char</a> | <a href="#little_u_value" class="noline">little_u_value</a> | <a href="#big_u_value" class="noline">big_u_value</a> | <a href="#escaped_char" class="noline">escaped_char</a> .
<a id="byte_value">byte_value</a>       = <a href="#octal_byte_value" class="noline">octal_byte_value</a> | <a href="#hex_byte_value" class="noline">hex_byte_value</a> .
<a id="octal_byte_value">octal_byte_value</a> = `\` <a href="#octal_digit" class="noline">octal_digit</a> <a href="#octal_digit" class="noline">octal_digit</a> <a href="#octal_digit" class="noline">octal_digit</a> .
<a id="hex_byte_value">hex_byte_value</a>   = `\` "x" <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> .
<a id="little_u_value">little_u_value</a>   = `\` "u" <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> .
<a id="big_u_value">big_u_value</a>      = `\` "U" <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a>
                           <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> .
<a id="escaped_char">escaped_char</a>     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
</pre>

``` go
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // rune literal containing single quote character
'aa'         // illegal: too many characters
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
```

### String literals ### {#String_literals}

A string literal represents a [string constant](#Constants) obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

Raw string literals are character sequences between back quotes, as in ``foo``. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters ('\r') inside raw string literals are discarded from the raw string value.

Interpreted string literals are character sequences between double quotes, as in `"bar"`. Within the quotes, any character may appear except newline and unescaped double quote. The text between the quotes forms the value of the literal, with backslash escapes interpreted as they are in [rune literals](#Rune_literals) (except that `\'` is illegal and `\"` is legal), with the same restrictions. The three-digit octal (`\`_nnn_) and two-digit hexadecimal (`\x`_nn_) escapes represent individual _bytes_ of the resulting string; all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual _characters_. Thus inside a string literal `\377` and `\xFF` represent a single byte of value `0xFF`=255, while `ÿ`, `\u00FF`, `\U000000FF` and `\xc3\xbf` represent the two bytes `0xc3` `0xbf` of the UTF-8 encoding of character U+00FF.

<pre class="ebnf"><a id="string_lit">string_lit</a>             = <a href="#raw_string_lit" class="noline">raw_string_lit</a> | <a href="#interpreted_string_lit" class="noline">interpreted_string_lit</a> .
<a id="raw_string_lit">raw_string_lit</a>         = "`" { <a href="#unicode_char" class="noline">unicode_char</a> | <a href="#newline" class="noline">newline</a> } "`" .
<a id="interpreted_string_lit">interpreted_string_lit</a> = `"` { <a href="#unicode_value" class="noline">unicode_value</a> | <a href="#byte_value" class="noline">byte_value</a> } `"` .
</pre>

``` go
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:

``` go
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

If the source code represents a character as two code points, such as a combining form involving an accent and a letter, the result will be an error if placed in a rune literal (it is not a single code point), and will appear as two code points if placed in a string literal.
