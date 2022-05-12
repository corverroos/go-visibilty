# go-visibility

This repo illustrates an approach to structuring Go code and how to reason about
identifier visibility and accessibility.

It proposes the **private-is-private-public-is-public** convention:
- Exported identifiers indicate **public-safe-to-use**; both by other packages (as per Go spec) AND other entities in the same package (by convention).
- Unexported identifiers indicate **private-not-safe-to-use**; both by other packages (enforced by Go spec) AND other entities in the same package (by convention).
- Types, interfaces, and functions should be unexported by default, ensuring "shy" code that is well isolated with small API surface and no leaking of internals.
- Accessibility/visibility of fields and methods is practically dictated by their type/interface so should follow the _public-safe-to-use_ vs _private-not-safe-to-use_ convention.
- This introduces additional semantic nuance to accessibility between identifiers _within_ the same package.

> The _private-is-private-public-is-public_ convention is contrasted again another 
> common approach of making fields and methods unexported by default (in addition to types and interfaces). 
> This goes against the _private-is-private-public-is-public_ convention
> since types within the same package access each other private identifiers and doesn't
> provide any indication of accessibility within the same package.

See [foo/foo.go](foo/foo.go) and other files for an illustration of this and examples.
 
## Go language spec

The following aspects of the Go language [spec](https://go.dev/ref/spec) is important to note in this context.

Go [identifiers](https://go.dev/ref/spec#Identifiers) are program entities that include:
- functions: `func foo(int) error {..}`
- types: `type foo struct{..}` or `type foo int`
- interfaces: `type foo interface {..}`
- methods: `func (b bar) foo() {..}`
- fields: `type bar struct{ foo int }`
- variables: `var foo int`
- constants: `const foo string`

Go source code is written in `*.go` files grouped by folders that construct [packages](https://go.dev/ref/spec#Packages):
- `*.go`: Source code included in the package (`package foo`) and compiled into program binary. All identifiers can access all other identifiers of the same package.
- `*_test.go`: Test code that can be executed with `go test` to test the package. Not compiled into program binary.
  - `*_internal_test.go`: Convention indicating _package internal_ test code (`package foo`) that can access all identifiers of the package under test.
  - `*_test.go`: Convention indicating _package external_ test code (`package foo_test`) that must explicitly import the package under test and can only access exported identifiers.

> Test code defined in `*_test.go` files (both _package internal_ and _package external_) cannot be imported
> by other packages; neither by other package's program source code nor by other package's test code.

Go programs are created by a [main package](https://go.dev/ref/spec#Program_execution):
- `package main`: Main package consisting of one or more source files.
- `func main() {..}`: Program entrypoint
- `import (..)`: Imports other packages included in the program, transitively.

> Main packages cannot be imported by other packages. 
> Main packages therefore only support _package internal_ test code, 
> since _package external_ `package main_test` would not be able to import the `main` package.

Identifiers are [exported](https://go.dev/ref/spec#Exported_identifiers) if 
the first character of the identifier name is an uppercase letter.
Exported identifiers can in general by accessed by other packages, unless (as mentioned above) the
identifier is defined is a main package or in test code.
```go
type unexported struct {..}
func unexported() {..}

type Exported struct {..}
var Exported int
```




