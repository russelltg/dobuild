package main

import "github.com/russelltg/dobuild/libdobuild"

func main() {
	t := libdobuild.CreatePrinterTarget("Hello!")
	t2 := libdobuild.CreatePrinterTarget("Hello to you too!")

	t2.AddDependency(t)

	t2.BuildWithDeps()

}
