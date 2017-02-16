package main

import (
	"log"
	"net/http"
)

type Content struct {
	Title string
	Body  []byte
}

type Conf struct {
	Host           string
	Port           int
	Bind           string
	Static         string
	TemplateRoot   string
	HeaderTemplate string
	FooterTemplate string
}

var conf = &Conf{
	Host:           "0.0.0.0",
	Port:           8081,
	Bind:           "0.0.0.0:8081",
	Static:         "/home/ichavero/goStuff/Go/src/sotolitolabs/sereke/static",
	TemplateRoot:   "/home/ichavero/goStuff/Go/src/sotolitolabs/sereke/tpl",
	HeaderTemplate: "/home/ichavero/goStuff/Go/src/sotolitolabs/sereke/tpl/header.tpl",
	FooterTemplate: "/home/ichavero/goStuff/Go/src/sotolitolabs/sereke/tpl/footer.tpl",
}

var Templates map[string]string

//TODO add command line options
/*func Init() {
	var static string
	Host:         "0.0.0.0",
	Port:         8081,
	Bind:         "0.0.0.0:8081",
	Static: "/home/ichavero/goStuff/Go/src/sotolitolabs/sereque/static",
	TemplateRoot: "/home/ichavero/goStuff/Go/src/sotolitolabs/sereque/tpl",

    flag.StringVar(&dir, "dir", conf.Static, "the directory to serve files from")
    flag.Parse()

}*/

func main() {
	Templates = make(map[string]string)
	router := NewRouter()
	//Add route for static content (might not be needed if behind a proxy)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir(conf.Static))))
	log.Fatal(http.ListenAndServe(conf.Bind, router))
}
