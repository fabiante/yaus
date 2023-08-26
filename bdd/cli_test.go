package bdd

import (
	"github.com/fabiante/yaus/bdd/drivers"
	"github.com/fabiante/yaus/bdd/specs"
	"testing"
)

func TestCLI(t *testing.T) {
	driver := drivers.NewCLIDriver("../cli/main.go")

	specs.TestAll(t, driver)
}
