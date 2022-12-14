package main

import (
	"flag"
	"fmt"
	"log"
	"path"

	"github.com/galavant-web-design/cardinal/pkg/generator"
	"github.com/galavant-web-design/cardinal/pkg/server"
)

func main() {
	buildOnly := flag.Bool("build", false, "only build files (don't serve and watch)")
	flag.Parse()

	rootPath := "."
	buildPath := path.Join(rootPath, "build")
	finder := generator.NewFinder(rootPath)
	gen := generator.Generator{BuildPath: buildPath}

	buildSite(finder, gen)

	if *buildOnly {
		fmt.Println("Done building")
		return
	}

	errorChannel := make(chan error)
	watcher, err := generator.Watch(rootPath, buildPath, errorChannel, func() { buildSite(finder, gen) })
	if err != nil {
		log.Fatalf("Unable to watch: %s", err)
	}
	defer watcher.Close()

	staticServer := server.Serve(gen.BuildPath, errorChannel)

	err = <-errorChannel
	defer staticServer.Close()
	log.Fatalf("Fatal error: %s", err)
}

func buildSite(finder generator.Finder, gen generator.Generator) {
	gen.ClearBuildDirectory()
	sourceFiles := finder.FindSourceFiles()
	gen.Build(sourceFiles, finder.TemplatePath)
}
