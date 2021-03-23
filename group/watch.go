package group

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Read group file from path and return Group object
func readGroupFile(name string, path string) (group Group, err error) {
	log.Printf("Loading %v", path)
	group.Name = name

	// Read group description from JSON
	jsonFile, err := os.Open(path)
	if err != nil {
		return group, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return group, err
	}
	if err = json.Unmarshal(byteValue, &group.Description); err != nil {
		return group, err
	}

	// Add modification time
	fileInfo, err := os.Stat(path)
	if err != nil {
		return group, err
	}
	group.Description.modTime = fileInfo.ModTime()
	return group, nil
}

// Reload group directory and scan for changed elements
// FIXME: scan removed files
func reloadDirectory() error {
	err := filepath.Walk(
		Directory,
		func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				log.Printf("Error when loading group file %v: %s", path, err)
				return nil
			}

			// Check file is JSON
			if fi.IsDir() {
				return nil
			}
			filename, err := filepath.Rel(Directory, path)
			if err != nil {
				log.Printf("Error when loading group file %v: %s", path, err)
				return nil
			}
			if !strings.HasSuffix(filename, ".json") {
				log.Printf("Unexpected extension for group file %v", path)
				return nil
			}

			// Check if group exists
			name := strings.TrimSuffix(filename, ".json")
			group := Group{}
			groupIndex := -1
			groupsMutex.Lock()
			defer groupsMutex.Unlock()
			for i, g := range Groups {
				if g.Name == name {
					group = g
					groupIndex = i
					break
				}
			}

			// Add if missing
			if groupIndex == -1 {
				group, err = readGroupFile(name, path)
				if err != nil {
					log.Printf("Error when opening group file %v: %s", path, err)
					return nil
				}
				Groups = append(Groups, group)
				return nil
			}

			// Update if file is newer
			if !group.Description.modTime.Before(fi.ModTime()) {
				return nil
			}
			group, err = readGroupFile(name, path)
			if err != nil {
				log.Printf("Error when opening group file %v: %s", path, err)
				return nil
			}
			Groups[groupIndex] = group
			return nil
		},
	)
	return err
}

// WatchGroups updates Groups variable when JSON changes
// We do not use fsnotify as this does not work on all filesystems.
func WatchGroups() {
	for {
		select {
		case <-time.After(time.Second * 3):
			if err := reloadDirectory(); err != nil {
				log.Printf("Failed to watch groups: %s", err)
				os.Exit(1)
			}
		}
	}
}
