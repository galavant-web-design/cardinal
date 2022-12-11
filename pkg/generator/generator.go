package generator

import (
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Generator struct {
	rootPath     string
	templatePath string
	BuildPath    string
}

type SourceFile struct {
	Path     string
	FileInfo os.FileInfo
}

func New(rootPath string) Generator {
	return Generator{
		rootPath:     rootPath,
		templatePath: path.Join(rootPath, "template.html"),
		BuildPath:    path.Join(rootPath, "build"),
	}
}

func (g Generator) ClearBuildDirectory() {
	err := os.RemoveAll(g.BuildPath)
	if err != nil {
		log.Fatalf("Unable to delete build directory: %s", err)
	}
}

func (g Generator) FindSourceFiles() []SourceFile {
	sourceFiles := make([]SourceFile, 0)

	err := filepath.Walk(g.rootPath, func(path string, info os.FileInfo, err error) error {
		if path == g.templatePath {
			return nil
		}

		if strings.HasPrefix(info.Name(), ".") && path != g.rootPath {
			if info.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		sourceFiles = append(sourceFiles, SourceFile{Path: path, FileInfo: info})
		return nil
	})
	if err != nil {
		log.Fatalf("Unable find source files: %s", err)
	}

	return sourceFiles
}

func (g Generator) Build(sourceFiles []SourceFile) {
	templateBytes := g.readTemplate()

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

func (g Generator) readTemplate() []byte {
	templateBytes, err := os.ReadFile(g.templatePath)
	if err != nil {
		log.Fatalf("Unable to read template %s: %s", g.templatePath, err)
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
