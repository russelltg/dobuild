// libdobuild has the library component of the dobuild library. It defines the
// language and the default targets
package dob

import "fmt"

// FileWatcherTarget Implements Target
type PrinterTarget struct {
	ToPrint string

	deps []Target
}

func (t *PrinterTarget) GetDeps() []Target {
	return t.deps
}
func (t *PrinterTarget) AddDependencies(toAdd []Target) {
	t.deps = append(t.deps, toAdd...)
}

func (t *PrinterTarget) Build() bool {
	fmt.Print(t.ToPrint)

	return true
}

func (t *PrinterTarget) NeedsRebuild() bool {

	return true
}

// CreatePrinterTarget creates a target that prints every time it is built, and is never out of date.
func NewPrinterTarget(toSay string) *PrinterTarget {
	return &PrinterTarget{ToPrint: toSay}
}
