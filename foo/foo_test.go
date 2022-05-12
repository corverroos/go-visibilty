package foo_test // <-- Note the package name

import (
	"github.com/corverroos/visibility/foo"
	"testing"
)

// This is `package external` test code.
// It is a test code since the file has `*_test.go` suffix.
// It is package external test code since it is in the `foo_test` package.
// It needs to explicitly import `package foo` to test it.
// By convention, package external test code is named `*_test.go`
// Other packages cannot import anything from test code.

// TestExternal is a `package external` test that can only access exported identifiers of foo.
func TestExternal(t *testing.T) {
	// var zero qbft.private  <- cannot instantiate unexported type.

	// It can access exported identifier.
	if foo.Public() != 0 {
		t.Fail()
	}

	// Note that Bad leaks the private unexported type (which is generally an anti-pattern),
	// and it's exported fields/methods is now accessible,
	// but since we follow private-is-private-public-is-public convention, this isn't a problem.
	if foo.Bad().PublicField != 0 { //
		t.Fail()
	}

	surprise(t)
}

// Exported is an exported type, but cannot be imported by other packages, since this is test code.
// This doesn't clash with Exported type defined in `foo_internal_test.go` since
// it is different packages.
type Exported struct{}

func surprise(t *testing.T) {
	// Surprisingly `public external` test code can import
	// `package internal` test code's exported identifiers.

	foo.TestInternal(t)
	var _ foo.TestExported

	// But once again, since we follow private-is-private-public-is-public, this
	// isn't a problem.
}
