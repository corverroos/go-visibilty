# go-visibility

This repo illustrates an approach to structuring Go code and how to reason about
identifier visibility and accessibility.

It proposes the **private-is-private-public-is-public** convention which is defined as:
- Exported identifiers indicate **public-safe-to-use**; both by other packages (as per Go spec) AND other entities in the same package (by convention).
- Unexported identifiers indicate **private-not-safe-to-use**; both by other packages (enforced by Go spec) AND other entities in the same package (by convention).
- Types, interfaces, and functions should be unexported by default, ensuring "shy" code that is well isolated with small API surface and no* leaking of internals.
- Accessibility/visibility of fields and methods is however practically dictated by their type/interface so are free follow the more nuanced _public-safe-to-use_ vs _private-not-safe-to-use_ convention.
- This introduces additional semantic nuance to accessibility between identifiers _within_ the same package.

> The main aim of this convention is to introduce intra-package semantics for accessibility of
> unexported types of the same package. Unexported types can therefore provide an indication 
> of their safe-to-use APIs to other types in the same package.
> 
> The point is that if a private type is not somehow returned outside the package, then 
> public fields/methods can be used to convey semantic meaning of a safe-to-use intra-package API.

It can be illustrated as follows:
```
                    Types                                Interface

              public  │ private                    public  │ private
            ┌─────────┼──────────┐               ┌─────────┼──────────┐
            │ Anyone  │ Other    │               │ Anyone  │ Other    │
      public│ in      │ types in │         public│ in      │ types in │
            │ world   │ package  │               │ world   │ package  │
Fields:  ───┼─────────┼──────────┤   Methods: ───┼─────────┼──────────┤
            │         │          │               │         │          │
     private│ No-one  │ No-one   │        private│ Anit-   │ Anti-    │
            │ else    │ else     │               │ pattern │ pattern  │
            │         │          │               │         │          │
            └─────────┴──────────┘               └─────────┴──────────┘
```

> The _private-is-private-public-is-public_ convention is contrasted against another 
> common approach of _also_ making fields and methods unexported by default (in addition to types and interfaces). 
> This goes against the _private-is-private-public-is-public_ convention
> since this leads to types within the same package accessing each other's private identifiers and doesn't
> provide any indication of accessibility within the same package. 
> 
> This "everything private by default" is a good approach, but it is limiting and 
> looses out on the opportunity for a more nuanced (but still safe) approach.
> Throwing out the baby with the bath water.

As with any convention, there are exceptions to the rule. If an exported type needs to provide an "internal" API
for other types in its own package, making that exported obviously leaks the internal API. Providing
that internal API via unexported identifiers is fine in that case.

Note that although exported fields and methods of unexported types can _technically_ be accessed by other packages
if the unexported type is returned by an exported function (or leaked in some other way).
This is generally not a concern since the _private-is-private-public-is-public_ convention states that public identifiers are safe-to-use.
This is also a derived second order risk, which is acceptable as a tradeoff for the improved semantic nuance introduced by the convention.

See [foo/foo.go](foo/foo.go) and other files for an illustration of this approach.


