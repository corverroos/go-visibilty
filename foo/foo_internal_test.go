package foo // <-- Note the package name

import "testing"

// This is `package internal` test code.
// It is a test code since the file has `*_test.go` suffix.
// It is package internal test code since it is in the `foo` package.
// By convention, package internal test code is named `*_internal_test.go`
// Other packages cannot import anything from test code.

// TestInternal is a `package internal` test that can access all identifiers (both exported and unexported)
func TestInternal(t *testing.T) {
	var zero private
	if Bad() != zero {
		t.Fail()
	}

	if Public() != 0 {
		t.Fail()
	}
}

// TestExported is an exported type from test code, so cannot be
// imported by other packages.
type TestExported struct {
	Field1 int
}
