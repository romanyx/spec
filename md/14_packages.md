## Packages ## {#Packages}

Go programs are constructed by linking together _packages_. A package in turn is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package. Those elements may be [exported](#Exported_identifiers) and used in another package.

### Source file organization ### {#Source_file_organization}

Each source file consists of a package clause defining the package to which it belongs, followed by a possibly empty set of import declarations that declare packages whose contents it wishes to use, followed by a possibly empty set of declarations of functions, types, variables, and constants.

<pre class="ebnf"><a id="SourceFile">SourceFile</a>       = [PackageClause](#PackageClause) ";" { [ImportDecl](#ImportDecl) ";" } { [TopLevelDecl](#TopLevelDecl) ";" } .
</pre>

### Package clause ### {#Package_clause}

A package clause begins each source file and defines the package to which the file belongs.

<pre class="ebnf"><a id="PackageClause">PackageClause</a>  = "package" [PackageName](#PackageName) .
<a id="PackageName">PackageName</a>    = [identifier](#identifier) .
</pre>

The PackageName must not be the [blank identifier](#Blank_identifier).

``` go
package math
```

A set of files sharing the same PackageName form the implementation of a package. An implementation may require that all source files for a package inhabit the same directory.

### Import declarations ### {#Import_declarations}

An import declaration states that the source file containing the declaration depends on functionality of the _imported_ package ([§Program initialization and execution](#Program_initialization_and_execution)) and enables access to [exported](#Exported_identifiers) identifiers of that package. The import names an identifier (PackageName) to be used for access and an ImportPath that specifies the package to be imported.

<pre class="ebnf"><a id="ImportDecl">ImportDecl</a>       = "import" ( <a href="#ImportSpec" class="noline">ImportSpec</a> | "(" { <a href="#ImportSpec" class="noline">ImportSpec</a> ";" } ")" ) .
<a id="ImportSpec">ImportSpec</a>       = [ "." | <a href="#PackageName" class="noline">PackageName</a> ] <a href="#ImportPath" class="noline">ImportPath</a> .
<a id="ImportPath">ImportPath</a>       = <a href="#string_lit" class="noline">string_lit</a> .
</pre>

The PackageName is used in [qualified identifiers](#Qualified_identifiers) to access exported identifiers of the package within the importing source file. It is declared in the [file block](#Blocks). If the PackageName is omitted, it defaults to the identifier specified in the [package clause](#Package_clause) of the imported package. If an explicit period (`.`) appears instead of a name, all the package's exported identifiers declared in that package's [package block](#Blocks) will be declared in the importing source file's file block and must be accessed without a qualifier.

The interpretation of the ImportPath is implementation-dependent but it is typically a substring of the full file name of the compiled package and may be relative to a repository of installed packages.

Implementation restriction: A compiler may restrict ImportPaths to non-empty strings using only characters belonging to [Unicode's](http://www.unicode.org/versions/Unicode6.3.0/) L, M, N, P, and S general categories (the Graphic characters without spaces) and may also exclude the characters `!"#$%&'()*,:;<=>?[\]^`{|}` and the Unicode replacement character U+FFFD.

Assume we have compiled a package containing the package clause `package math`, which exports function `Sin`, and installed the compiled package in the file identified by `"lib/math"`. This table illustrates how `Sin` is accessed in files that import the package after the various types of import declaration.

``` go
Import declaration          Local name of Sin

import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin
```

An import declaration declares a dependency relation between the importing and imported package. It is illegal for a package to import itself, directly or indirectly, or to directly import a package without referring to any of its exported identifiers. To import a package solely for its side-effects (initialization), use the [blank](#Blank_identifier) identifier as explicit package name:

``` go
import _ "lib/math"
```

### An example package ### {#An_example_package}

Here is a complete Go package that implements a concurrent prime sieve.

``` go
package main

import "fmt"

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i  // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {  // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i  // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int)  // Create a new channel.
	go generate(ch)       // Start generate() as a subprocess.
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
```
