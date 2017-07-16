## Представление исходного кода ## {#Source_code_representation}

Исходный код представляет собой символы Юникод, закодированные в кодировке [UTF-8](http://en.wikipedia.org/wiki/UTF-8). Текст не нормализован, то-есть монолитный символ из одной кодовой точки будет отличаться от идентичного созданного посредством соединения двух кодовых точек, состоящих из базового символа и модифицирующего символа. Для простоты, этот документ будет использовать неквалифицированный термин _символ_ для обозначения кодовой точки Unicode в исходном коде

Каждая кодовая точка уникальна, то-есть символы в верхнем и нижнем регистре отдельные символы

Ограничение: Для совместимости с другими инструментами, компилятор может отклонить символ NUL (U+0000) в исходном коде

Ограничение: Для совместимости с другими инструментами, компилятор может отклонить символ метки последовательности байтов (U+FEFF) если он установлен в начале исходного кода. Также маркер последовательности байтов может быть проигнорирован где-либо еще в исходном коде.

### Символы ### {#Characters}

Следующие определения используются для обозначения конкретных классов символов Unicode

<pre class="ebnf"><a id="newline">newline</a>        = /* the Unicode code point U+000A */ .
<a id="unicode_char">unicode_char</a>   = /* an arbitrary Unicode code point except newline */ .
<a id="unicode_letter">unicode_letter</a> = /* a Unicode code point classified as "Letter" */ .
<a id="unicode_digit">unicode_digit</a>  = /* a Unicode code point classified as "Number, decimal digit" */ .
</pre>

В стандарте [Unicode Standard 8.0](http://www.unicode.org/versions/Unicode8.0.0/), секция 4.5 "Общие категории" определены категории символов. Go воспринимает все символы из Letter категорий Lu, Ll, Lt, Lm и Lo как строковые символы Unicode и те которые относятся к категории Number Nd как цифры Unicode.

### Буквы и цифры ### {#Letters_and_digits}

Символ подчеркивания `_` (U+005F) воспринимается как буква.

<pre class="ebnf"><a id="letter">letter</a>        = <a href="#unicode_letter" class="noline">unicode_letter</a> | "_" .
<a id="decimal_digit">decimal_digit</a> = "0" … "9" .
<a id="octal_digit">octal_digit</a>   = "0" … "7" .
<a id="hex_digit">hex_digit</a>     = "0" … "9" | "A" … "F" | "a" … "f" .
</pre>

