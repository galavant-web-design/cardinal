package generator

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Finder struct {
	rootPath     string
	TemplatePath string
}

type SourceFile struct {
	Path     string
	FileInfo os.FileInfo
}

func NewFinder(rootPath string) Finder {
	return Finder{
		rootPath:     rootPath,
		TemplatePath: path.Join(rootPath, "template.html"),
	}
}

func (f Finder) FindSourceFiles() []SourceFile {
	sourceFiles := make([]SourceFile, 0)

	err := filepath.Walk(f.rootPath, func(path string, info os.FileInfo, err error) error {
		if path == f.TemplatePath {
			return nil
		}

		if strings.HasPrefix(info.Name(), ".") && path != f.rootPath {
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
