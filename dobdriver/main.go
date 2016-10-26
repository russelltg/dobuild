package main

import (
	"github.com/russelltg/dobuild/dob"
)

func main() {
	t := dob.NewPrinterTarget("Hello!")
	t2 := dob.NewPrinterTarget("Hello to you too!")

	t2.AddDependencies([]dob.Target{t})

	dob.BuildTarget(t2)

}
