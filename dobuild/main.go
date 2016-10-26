package main

import "github.com/russelltg/dobuild/libdobuild"

func main() {
	t := libdobuild.NewPrinterTarget("Hello!")
	t2 := libdobuild.NewPrinterTarget("Hello to you too!")

	t2.AddDependencies([]libdobuild.Target{t})

	libdobuild.BuildTarget(t2)

}
