package generator

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func Watch(rootPath string, buildPath string, action func()) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if strings.HasSuffix(event.Name, "~") || strings.HasPrefix(event.Name, buildPath) {
					continue
				}

				action()
			case watchError, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error: ", watchError)
			}
		}
	}()

	return watcher, watcher.Add(rootPath)
}
