package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderLayout prints the entire layout
func RenderLayout(c *Content, w http.ResponseWriter) {
	t, err := template.ParseFiles(conf.HeaderTemplate,
		conf.FooterTemplate,
		conf.TemplateRoot+"/main.tpl",
		conf.TemplateRoot+"/messages.tpl",
		conf.TemplateRoot+"/message.tpl",
		conf.TemplateRoot+"/folders.tpl")
	if err != nil {
		log.Fatalf("Can't render template %s: %s", t, err.Error())
	}
	t.ExecuteTemplate(w, "header", c)
	t.ExecuteTemplate(w, "main", c)
	//t.ExecuteTemplate(w, "folders", c)
	//t.ExecuteTemplate(w, "messages", c)
	t.ExecuteTemplate(w, "footer", c)
}

// Render prints a single template
// TODO refactor
func Render(path string, name string, content interface{}, w http.ResponseWriter) {
	log.Printf("Executing template[%s] : %s\n", name, path)
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalf("Can't render template %s: %s", t, err.Error())
	}
	log.Printf("Render::content: %s", content)
	/*for elem := range content {
		log.Printf("RENDER::Name: %s\n", elem.Name)
	}*/
	t.ExecuteTemplate(w, name, content)
}

// RenderMultiple prints a single template or a group of templates
// TODO refactor
func RenderMultiple(templates map[string]string, c *Content, w http.ResponseWriter) {
	for k, v := range templates {
		log.Printf("Executing template[%s] : %s\n", k, v)
		t, err := template.ParseFiles(v)
		if err != nil {
			log.Fatalf("Can't render template %s: %s", t, err.Error())
		}
		t.ExecuteTemplate(w, k, c)
	}
}

func TestRenderLayout(c *Content, w http.ResponseWriter) {
	t, err := template.ParseFiles(conf.HeaderTemplate,
		conf.FooterTemplate,
		conf.TemplateRoot+"/main.tpl",
		conf.TemplateRoot+"/messages.tpl",
		conf.TemplateRoot+"/message.tpl",
		conf.TemplateRoot+"/folders.tpl")
	if err != nil {
		log.Fatalf("Can't render template %s: %s", t, err.Error())
	}
	t.ExecuteTemplate(w, "header", c)
	t.ExecuteTemplate(w, "main", c)
	//t.ExecuteTemplate(w, "messages", c)
	t.ExecuteTemplate(w, "footer", c)
}
