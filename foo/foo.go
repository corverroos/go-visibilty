package foo

// Public is a public-safe-to-use exported function of foo package.
func Public() int {
	// Can instantiate and do things with private struct types.
	var p private

	// Should however only access public-safe-to-se API of
	// other types (even if private fields are technically accessible).
	return p.PublicMethod() * p.PublicField
}

// private is unexported, so only accessible to other types in this package.
type private struct {
	// privateField is unexported so indicates private-not-safe-to-use by all other types.
	privateField int

	// PublicField is exported so indicates public-safe-to-use by all other types.
	PublicField int
}

// PublicMethod is exported so indicates public-safe-to-use by other types.
func (p private) PublicMethod() int {
	// types can access their own private methods.
	return p.privateMethod()
}

// privateMethod is unexported so indicates private-not-safe-to-use by other types.
func (p private) privateMethod() int {
	// types can access their own private fields.
	return p.privateField
}

// Bad is generally an anti-pattern since it "leaks" an unexported type.
//
// Note that private's PublicMethod and PublicField will now be accessible to other packages,
// but since we follow public-safe-to-use convention, this isn't a concern.
//
// Note there are edge use-cases where leaking an unexported type is applicable though, e.g. functional options.
func Bad() private {
	return private{}
}
