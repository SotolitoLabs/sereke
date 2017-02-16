package main

import "net/http"

// Route defines the routes and their elements
type Route struct {
	Name        string
	Method      string
	Pattern     string
	Headers     []string
	HandlerFunc http.HandlerFunc
}

//Routes is the list of routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		[]string{"Content-type", "(text/html*|application/json)"},
		Index,
	},
	Route{
		"Login",
		"GET",
		"/login",
		[]string{"Content-type", "(text/html*|application/json)"},
		Login,
	},
	Route{
		"Folders",
		"GET",
		"/folders",
		[]string{"Content-type", "(text/html*|application/json)"},
		ShowFolders,
	},
	Route{
		"Folder",
		"GET",
		"/folder/{folderId}",
		[]string{"Content-type", "(text/html*|application/json)"},
		ShowFolder,
	},
	Route{
		"Messages",
		"GET",
		"/messages",
		[]string{"Content-type", "(text/html*|application/json)"},
		ShowMessages,
	},
}
