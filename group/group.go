package group

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sync"
	"time"
)

// Directory containing groups JSON description
var Directory string

// Description from github.com/jech/galene/group
type Description struct {
	fileName       string        `json:"-"`
	modTime        time.Time     `json:"-"`
	fileSize       int64         `json:"-"`
	Description    string        `json:"description,omitempty"`
	Contact        string        `json:"contact,omitempty"`
	Comment        string        `json:"comment,omitempty"`
	Redirect       string        `json:"redirect,omitempty"`
	Public         bool          `json:"public,omitempty"`
	MaxClients     int           `json:"max-clients,omitempty"`
	MaxHistoryAge  int           `json:"max-history-age,omitempty"`
	AllowAnonymous bool          `json:"allow-anonymous,omitempty"`
	AllowRecording bool          `json:"allow-recording,omitempty"`
	AllowSubgroups bool          `json:"allow-subgroups,omitempty"`
	Autolock       bool          `json:"autolock,omitempty"`
	Autokick       bool          `json:"autokick,omitempty"`
	Op             []interface{} `json:"op,omitempty"`
	Presenter      []interface{} `json:"presenter,omitempty"`
	Other          []interface{} `json:"other,omitempty"`
	Codecs         []string      `json:"codecs,omitempty"`
}

// Group contains a name and a description
type Group struct {
	Name        string
	Description Description
}

// Groups contains all groups
var Groups []Group = make([]Group, 0)
var groupsMutex sync.Mutex

// WatchGroups updates Groups variable when JSON changes
func WatchGroups() {
	// TODO
}

// Create new group
func Create(newGroup Group) error {
	// Verify new group name is a slug
	matched, err := regexp.MatchString(`^[a-z0-9_-]+$`, newGroup.Name)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("name is not a slug")
	}

	// Check group does not already exist
	groupsMutex.Lock()
	defer groupsMutex.Unlock()
	for _, g := range Groups {
		if g.Name == newGroup.Name {
			return errors.New("name already exist")
		}
	}

	// Create and write JSON
	content, err := json.MarshalIndent(newGroup, "", "  ")
	if err != nil {
		return err
	}
	filePath := path.Join(Directory, newGroup.Name+".json")
	if err = ioutil.WriteFile(filePath, content, 0644); err != nil {
		return err
	}

	Groups = append(Groups, newGroup)
	return nil
}

// Update group
func Update(newGroup Group) error {
	groupsMutex.Lock()
	defer groupsMutex.Unlock()

	for i, g := range Groups {
		if g.Name == newGroup.Name {

			// Create and write JSON
			content, err := json.MarshalIndent(newGroup, "", "  ")
			if err != nil {
				return err
			}
			filePath := path.Join(Directory, g.Name+".json")
			if err = ioutil.WriteFile(filePath, content, 0644); err != nil {
				return err
			}

			Groups[i] = newGroup
			return nil
		}
	}
	return errors.New("not found")
}

// Delete group by name
func Delete(name string) error {
	groupsMutex.Lock()
	defer groupsMutex.Unlock()

	for i, g := range Groups {
		if g.Name == name {
			// Remove JSON
			filePath := path.Join(Directory, g.Name+".json")
			if err := os.Remove(filePath); err != nil {
				return err
			}

			Groups = append(Groups[:i], Groups[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
