package libdobuild

import "testing"

func TestCompileDependencies(t *testing.T) {
	parent := NewPrinterTarget("Hello")

	c1 := NewPrinterTarget("2")
	c2 := NewPrinterTarget("3")

	parent.AddDependencies([]Target{c1, c2})

	BuildTarget(parent)
}
