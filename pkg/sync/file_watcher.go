package sync

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

type FileWatcher struct {
	Changes chan ChangeEvent
}

type ChangeEvent int

const (
	FileChanged ChangeEvent = iota
)

func NewFileWatcher(filePath string) *FileWatcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(fmt.Errorf("error creating file watcher: %v", err))
	}

	err = watcher.Add(filePath)
	if err != nil {
		panic(fmt.Errorf("error adding file to watcher: %v", err))
	}

	changes := make(chan ChangeEvent, 1)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					changes <- FileChanged
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Printf("Error watching file: %v\n", err)
			}
		}
	}()

	return &FileWatcher{
		Changes: changes,
	}
}

func (fw *FileWatcher) Watch() {
	// Watch for file changes and send events to the Changes channel
	select {}
}
