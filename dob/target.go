package dob

type Target interface {
	GetDeps() []Target
	AddDependencies([]Target)
	Build() bool
	NeedsRebuild() bool
}

func BuildTarget(t Target) {

	for _, dep := range t.GetDeps() {
		if dep.NeedsRebuild() {
			res := dep.Build()
			if !res {
				// TODO: way way way better logging
				panic("Error building dependency!")
			}
		}
	}
	if t.NeedsRebuild() {
		t.Build()
	}
}
