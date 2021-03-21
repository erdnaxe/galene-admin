package webserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPAddr is listening address of the webserver
var HTTPAddr string

// StaticRoot is the path the /static folder served at the root of webserver
var StaticRoot string

// Serve webserver
func Serve() {
	// Set up router
	r := mux.NewRouter()
	r.Path("/api/group").Methods(http.MethodGet).HandlerFunc(listGroups)
	r.Path("/api/group").Methods(http.MethodPost).HandlerFunc(createGroup)
	r.Path("/api/group/{name}").Methods(http.MethodGet).HandlerFunc(readGroup)
	r.Path("/api/group/{name}").Methods(http.MethodPut).HandlerFunc(updateGroup)
	r.Path("/api/group/{name}").Methods(http.MethodDelete).HandlerFunc(deleteGroup)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(StaticRoot)))

	// HTTP server
	log.Printf("HTTP server listening on %s", HTTPAddr)
	log.Fatal(http.ListenAndServe(HTTPAddr, r))
}
