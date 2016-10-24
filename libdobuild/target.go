package libdobuild

type Target struct {
	Deps         []*Target
	Build        func(*Target) bool
	NeedsRebuild func(*Target) bool
}

func (t *Target) AddDependency(toAdd *Target) {
	t.Deps = append(t.Deps, toAdd)
}

func (t *Target) BuildWithDeps() {

	for _, dep := range t.Deps {
		if dep.NeedsRebuild(t) {
			res := dep.Build(t)
			if !res {
				// TODO: way way way better logging
				panic("Error building dependency!")
			}
		}
	}
	if t.NeedsRebuild(t) {
		t.Build(t)
	}
}
