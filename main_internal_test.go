package main

import "testing"

// TestInternal is a `package internal` test.
func TestInternal(t *testing.T) {
	// Access to everything in main, both exported and unexported.
	main()

	p := MakeParams()
	if p.P1 != 1 {
		t.Fail()
	}
}
