package main

import (
	"github.com/tygern/generator/internal/generator"
)

func main() {
	g := generator.New(".")

	g.ClearBuildDirectory()
	sourceFiles := g.FindSourceFiles()
	g.Build(sourceFiles)
}
