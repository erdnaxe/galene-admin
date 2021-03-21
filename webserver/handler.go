package webserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/erdnaxe/galene-admin/group"
	"github.com/gorilla/mux"
)

func listGroups(w http.ResponseWriter, r *http.Request) {
	// Try to encode group.Groups
	if err := json.NewEncoder(w).Encode(group.Groups); err != nil {
		http.Error(w, "Failed to encode JSON.", http.StatusInternalServerError)
		log.Printf("Failed to encode JSON: %s", err)
	}
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	// Read body content
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body.", http.StatusBadRequest)
		log.Printf("Failed to read body: %s", err)
		return
	}

	// Unmarshal and add
	var g group.Group
	if err = json.Unmarshal(reqBody, &g); err != nil {
		http.Error(w, "Failed to unmarshal JSON.", http.StatusBadRequest)
		log.Printf("Failed to unmarshal JSON: %s", err)
		return
	}
	if err = group.Create(g); err != nil {
		http.Error(w, "Failed to create group.", http.StatusInternalServerError)
		log.Printf("Failed to create group: %s", err)
		return
	}

	// Return new group
	if err = json.NewEncoder(w).Encode(g); err != nil {
		http.Error(w, "Failed to encode JSON.", http.StatusInternalServerError)
		log.Printf("Failed to encode JSON: %s", err)
	}
}

func readGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	for _, g := range group.Groups {
		if g.Name == name {
			json.NewEncoder(w).Encode(g)
			return
		}
	}

	http.Error(w, "Group not found.", http.StatusNotFound)
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Read body content
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body.", http.StatusBadRequest)
		log.Printf("Failed to read body: %s", err)
		return
	}
	var g group.Group
	if err = json.Unmarshal(reqBody, &g); err != nil {
		http.Error(w, "Failed to unmarshal JSON.", http.StatusBadRequest)
		log.Printf("Failed to unmarshal JSON: %s", err)
		return
	}

	// Check that group name is consistent
	if g.Name != vars["name"] {
		http.Error(w, "Group name does not match between query and JSON.", http.StatusBadRequest)
		return
	}

	// Update
	if err = group.Update(g); err != nil {
		http.Error(w, "Failed to update group.", http.StatusInternalServerError)
		log.Printf("Failed to update group: %s", err)
		return
	}

	// Return updated group
	if err = json.NewEncoder(w).Encode(g); err != nil {
		http.Error(w, "Failed to encode JSON.", http.StatusInternalServerError)
		log.Printf("Failed to encode JSON: %s", err)
	}
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if err := group.Delete(name); err != nil {
		http.Error(w, "Failed to delete group.", http.StatusInternalServerError)
		log.Printf("Failed to delete group: %s", err)
	}
}
