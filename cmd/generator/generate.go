package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SourceFile struct {
	Path     string
	FileInfo os.FileInfo
}

func main() {
	root := "."

	buildDir := path.Join(root, "build")
	err := os.RemoveAll(buildDir)
	if err != nil {
		log.Fatalf("Unable to delete build directory")
	}

	templatePath := path.Join(root, "template.html")

	sourceFiles := make([]SourceFile, 0)

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == templatePath {
			return nil
		}

		sourceFiles = append(sourceFiles, SourceFile{Path: path, FileInfo: info})
		return nil
	})

	templateBytes, _ := os.ReadFile(templatePath)

	for _, f := range sourceFiles {
		destination := path.Join(buildDir, f.Path)

		if f.FileInfo.IsDir() {
			_ = os.MkdirAll(destination, 0700)
			continue
		}

		if strings.HasSuffix(f.Path, ".html") {
			input, _ := os.ReadFile(f.Path)
			input = bytes.ReplaceAll(templateBytes, []byte("<#content/>"), input)
			err := os.WriteFile(destination, input, 0700)
			if err != nil {
				log.Fatalf("Unable to copy %s: %s", f.Path, err)
			}
		} else {
			from, err := os.Open(f.Path)
			if err != nil {
				log.Fatal(err)
			}
			defer from.Close()

			to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0700)
			if err != nil {
				log.Fatal(err)
			}
			defer to.Close()

			_, err = io.Copy(to, from)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
