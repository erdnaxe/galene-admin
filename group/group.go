package group

import (
	"errors"
	"time"
)

// Directory containing groups JSON description
var Directory string

// Group description from github.com/jech/galene/group
// fileName is changed to (Name + ".json")
type Group struct {
	Name           string        `json:"-"`
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

// Groups contains all groups
var Groups []Group

// WatchGroups updates Groups variable when JSON changes
func WatchGroups() {
	// TODO
}

// Create new group
func Create(g Group) error {
	Groups = append(Groups, g)

	// TODO: create JSON file
	return nil
}

// Update group
func Update(newGroup Group) error {
	for i, g := range Groups {
		if g.Name == newGroup.Name {
			Groups[i] = g
			// TODO: write JSON file
			return nil
		}
	}
	return errors.New("not found")
}

// Delete group by name
func Delete(name string) error {
	for i, g := range Groups {
		if g.Name == name {
			Groups = append(Groups[:i], Groups[i+1:]...)
			// TODO: remove JSON file
			return nil
		}
	}
	return errors.New("not found")
}
