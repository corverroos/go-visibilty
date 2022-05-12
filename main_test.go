package main_test

import "testing"

// TestExternal is a `package external` test, but since main package cannot be imported
// it cannot actually test anything in the main package.
func TestExternal(t *testing.T) {
	// No access to anything in main package.
}
