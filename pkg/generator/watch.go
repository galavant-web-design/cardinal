package generator

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func Watch(rootPath string, buildPath string, errorChannel chan error, action func()) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Println("watcher stopping")
					errorChannel <- errors.New("watcher event channel closed")
					return
				}
				if strings.HasSuffix(event.Name, "~") || strings.HasPrefix(event.Name, buildPath) {
					continue
				}

				action()
			case watchError, ok := <-watcher.Errors:
				if !ok {
					log.Println("watcher stopping")
					errorChannel <- errors.New("watcher error channel closed")
					return
				}
				log.Println("error: ", watchError)
			}
		}
	}()

	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			return nil
		}

		if info.Name() != rootPath && (strings.HasPrefix(info.Name(), ".") || path == buildPath) {
			fmt.Printf("Skipping %s\n", info.Name())
			return filepath.SkipDir
		}

		return watcher.Add(info.Name())
	})

	return watcher, err
}
