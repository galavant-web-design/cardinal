package main

import (
	"fmt"

	"github.com/tygern/cardinal/pkg/generator"
)

func main() {
	g := generator.New(".")

	g.ClearBuildDirectory()
	sourceFiles := g.FindSourceFiles()
	g.Build(sourceFiles)

	fmt.Println("Done!")
}
