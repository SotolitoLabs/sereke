package main

import (
	"fmt"
	"github.com/emersion/go-imap/client"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	//For testing:
	"github.com/emersion/go-imap"
)

//Manage sessions
var store = sessions.NewCookieStore([]byte("this_secret_should_be_configurable"))

type Connection struct {
	Sid  string
	Conn *client.Client
}

var Connections map[string]Connection

// Index returns the main page
func Index(w http.ResponseWriter, r *http.Request) {
	page := &Content{Title: "pagina principal",
		Body: []byte("contenido de página principal")}
	w.Header().Set("Content-type", "text/html")
	//RenderLayout(page, w)
	TestRenderLayout(page, w)
	fmt.Fprintln(w, "Welcome!")
}

// Login manages the authentication service
func Login(w http.ResponseWriter, r *http.Request) {
	//template := conf.TemplateRoot + "/login.tpl"
	c, _ := ImapLogin("mail.sotolitolabs.com", "993",
		"imcsk8@sotolitolabs.com", "c0m0lv1d4rl0")
	sid := "sereke-sid"
	/*session, err := store.Get(r, "session-name")
	  if err != nil {
	      http.Error(w, err.Error(), http.StatusInternalServerError)
	      return
	  } */
	log.Printf("Login :: Connection: %s", c)
	Connections[sid] = Connection{Sid: sid, Conn: c}
	//fmt.Fprintln(w, "Template : %s", template)
}

// ShowFolders gets the folder list
func ShowFolders(w http.ResponseWriter, r *http.Request) {
	/*page := &Content{Title: "Folders",
	Body: []byte("folder contents")}
	*/
	//Login(w, r) //Assert login
	//log.Printf("CONN?: %s", Connections["sereke-sid"].Conn)
	//messages := GetMailboxes(Connections["sereke-sid"].Conn)
	messages := PopulateMailBoxes()
	tpl := conf.TemplateRoot + "/folders.tpl"
	w.Header().Set("Content-type", "text/html")
	Render(tpl, "folders", messages, w)
	//fmt.Printf("ShowFolders :: Template : %s", tpl)
}

/*

  FOR TESTING PURPOUSES

*/

func PopulateMailBoxes() map[string]*imap.MailboxInfo {
	log.Printf("PopulateMailboxes....\n")
	folders := make(map[string]*imap.MailboxInfo)
	folders["INBOX"] = &imap.MailboxInfo{Name: "INBOX", Delimiter: "/"}
	folders["Sent"] = &imap.MailboxInfo{Name: "Sent", Delimiter: "/"}
	folders["Trash"] = &imap.MailboxInfo{Name: "Trash", Delimiter: "/"}
	folders["Spam"] = &imap.MailboxInfo{Name: "Spam", Delimiter: "/"}
	return folders
}

/*

  END TESTING  CODE

*/

// ShowMessages gets the message list
func ShowMessages(w http.ResponseWriter, r *http.Request) {
	page := &Content{Title: "Messages",
		Body: []byte("message list")}
	tpl := conf.TemplateRoot + "/messages.tpl"
	w.Header().Set("Content-type", "text/html")
	Render(tpl, "messages", page, w)
	fmt.Printf("ShowMessages :: Template : %s", tpl)
}

// ShowFolder gets the folder contents
func ShowFolder(w http.ResponseWriter, r *http.Request) {
	template := conf.TemplateRoot + "/message_list.tpl"
	fmt.Fprintln(w, "Template : %s", template)
}

/*
func serialize(data WebData) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
*/
