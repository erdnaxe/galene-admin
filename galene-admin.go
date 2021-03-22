package main

import (
	"flag"

	"github.com/erdnaxe/galene-admin/group"
	"github.com/erdnaxe/galene-admin/webserver"
)

func main() {
	// Parse command line arguments
	flag.StringVar(&webserver.HTTPAddr, "http", ":8080", "web server `address`")
	flag.StringVar(&webserver.StaticRoot, "static", "./static/",
		"web server root `directory`")
	flag.StringVar(&group.Directory, "groups", "./groups/",
		"group description `directory`")
	flag.Parse()

	// Start routines
	go group.WatchGroups()
	go webserver.Serve()

	// Wait for routines
	select {}
}
