package bar_test // <-- Note the package name

import (
	"github.com/corverroos/go-visibility/foo"
	"testing"
)

func TestInternal(t *testing.T) {
	// Foo's exported types are accessible
	foo.Public()
	foo.Bad().PublicMethod()

	// But nothing from the test code is accessible.
	// foo.TestInternal(t)
	// foo.TestExport

	// The external test package isn't importable.
	// foo_test.TestExternal(t)
	// foo_test.TestExport
}
