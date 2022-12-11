package generator

import (
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type Generator struct {
	BuildPath string
}

func (g Generator) ClearBuildDirectory() {
	err := os.RemoveAll(g.BuildPath)
	if err != nil {
		log.Fatalf("Unable to delete build directory: %s", err)
	}
}

func (g Generator) Build(sourceFiles []SourceFile, templatePath string) {
	templateBytes := g.readTemplate(templatePath)

	for _, sourceFile := range sourceFiles {
		destination := path.Join(g.BuildPath, sourceFile.Path)

		if sourceFile.FileInfo.IsDir() {
			makeDirectory(destination)
			continue
		}

		if strings.HasSuffix(sourceFile.Path, ".html") {
			applyTemplate(sourceFile, templateBytes, destination)
		} else {
			copyFile(sourceFile, destination)
		}
	}

}

func (g Generator) readTemplate(templatePath string) []byte {
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("Unable to read template %s: %s", templatePath, err)
	}
	return templateBytes
}

func makeDirectory(destination string) {
	err := os.MkdirAll(destination, 0700)
	if err != nil {
		log.Fatalf("Unable to create directory %s: %s", destination, err)
	}
}

func applyTemplate(sourceFile SourceFile, templateBytes []byte, destination string) {
	input, _ := os.ReadFile(sourceFile.Path)
	input = bytes.ReplaceAll(templateBytes, []byte("<#content/>"), input)
	err := os.WriteFile(destination, input, 0700)
	if err != nil {
		log.Fatalf("Unable to copy %s: %s", sourceFile.Path, err)
	}
}

func copyFile(sourceFile SourceFile, destination string) {
	from, err := os.Open(sourceFile.Path)
	if err != nil {
		log.Fatalf("Unable to open %s: %s", sourceFile.Path, err)
	}
	defer func(from *os.File) {
		_ = from.Close()
	}(from)

	to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0700)
	if err != nil {
		log.Fatalf("Unable to create %s: %s", destination, err)
	}
	defer func(to *os.File) {
		_ = to.Close()
	}(to)

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatalf("Unable to copy %s: %s", sourceFile.Path, err)
	}
}
