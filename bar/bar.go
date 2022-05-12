package bar

import "github.com/corverroos/visibility/foo"

func bar() {
	// Exported identifiers of foo is accessible.
	foo.Public()
	foo.Bad().PublicMethod()

	// Test code isn't accessible (neither package internal nor package external tests).
	// foo.TestExported
	// foo_test.TestExport
}
