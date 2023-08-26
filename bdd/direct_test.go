package bdd

import (
	"github.com/fabiante/yaus/app"
	"github.com/fabiante/yaus/bdd/drivers"
	"github.com/fabiante/yaus/bdd/specs"
	"testing"
)

func TestDirect(t *testing.T) {
	service := app.NewService()
	driver := drivers.NewDirectDriver(service)

	specs.TestAll(t, driver)
}
