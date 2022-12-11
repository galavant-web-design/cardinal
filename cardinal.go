package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tygern/cardinal/pkg/generator"
	"github.com/tygern/cardinal/pkg/server"
)

func main() {
	serve := flag.Bool("serve", false, "serve files")
	flag.Parse()

	g := generator.New(".")

	g.ClearBuildDirectory()
	sourceFiles := g.FindSourceFiles()
	g.Build(sourceFiles)
	fmt.Println("Done building")

	if !*serve {
		return
	}

	err := server.Serve(g.BuildPath)
	if err != nil {
		log.Fatal(err)
	}
}
