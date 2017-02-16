package main

import (
	"crypto/tls"
	"errors"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"io/ioutil"
	"log"
	//"mime"
	"net/mail"
)

// ImapLogin performs a login to a IMAP server
func ImapLogin(server string, port string, username string,
	password string) (*client.Client, error) {

	log.Println("Connecting to server...")

	//Allow any certificate
	tlsconf := &tls.Config{InsecureSkipVerify: true}
	// Connect to server
	c, err := client.DialTLS(server+":"+port, tlsconf)
	if err != nil {
		log.Printf("Error connecting to server %s:%s, %s", server, port, err)
		return nil, err
	}
	log.Println("Connected")

	// Login
	if err := c.Login(username, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	return c, nil
}

// GetMailboxes returns the mailboxes available for the account
func GetMailboxes(c *client.Client) map[string]*imap.MailboxInfo {
	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	go func() {
		// c.List will send mailboxes to the channel and close it when done
		if err := c.List("", "*", mailboxes); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Mailboxes:")
	folders := make(map[string]*imap.MailboxInfo)
	for m := range mailboxes {
		folders[m.Name] = m
		log.Printf("* %s", m.Name)
	}

	return folders
}

//GetMessageList returns a list of messages
func GetMessageList(start uint32, end uint32, mailbox string,
	c *client.Client) (map[uint32]*imap.Message, error) {

	mbox, err := c.Select(mailbox, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Flags for %s: %s", mailbox, mbox.Flags)

	// TODO make this configurable
	/*if mbox.Messages > 10 {
		from = mbox.Messages - 9
	}*/
	seqset := new(imap.SeqSet)
	//seqset.AddRange(from, to)
	seqset.AddRange(start, end)

	log.Printf("Seqset: %s", seqset.String())
	log.Printf("REAL MSG ID's: from: %d, to: %d, total: %d", start, end, mbox.Messages)

	messages := make(chan *imap.Message, 10)
	go func() {
		if err := c.Fetch(seqset, []string{imap.EnvelopeMsgAttr}, messages); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("messages from %d to %d:", start, end)

	msgList := make(map[uint32]*imap.Message)
	for msg := range messages {
		log.Printf("Adding %d", msg.SeqNum)
		msgList[msg.SeqNum] = msg
	}

	return msgList, nil

}

//GetMessage returns a message by ID
func GetMessage(messageID uint32, mailbox string,
	c *client.Client) (*mail.Message, error) {
	// Select INBOX
	if mailbox == "" {
		log.Printf("Missing mailbox")
		return nil, errors.New("Missing mailbox")
	}
	mbox, err := c.Select(mailbox, false)
	if err != nil {
		log.Printf("Error opening mailbox %s: %s", mailbox, err)
		return nil, err
	}

	// Get the last message
	if mbox.Messages == 0 {
		log.Println("No message in mailbox")
		return nil, errors.New("No messages in mailbox")
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(messageID, messageID)

	// Get the whole message body
	attrs := []string{"BODY[]"}
	messages := make(chan *imap.Message, 1)
	go func() {
		if err := c.Fetch(seqset, attrs, messages); err != nil {
			log.Printf("Couldn't get message %d: %s", messageID, err)
		}
	}()

	log.Println("message #%d:", messageID)
	msg := <-messages
	r := msg.GetBody("BODY[]")
	if r == nil {
		log.Printf("Server didn't returned message body for: %d", messageID)
	}

	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Printf("Error reading message %d: %s", messageID, err)
		return nil, err
	}

	return m, nil
}

//GetBody returns the body of a message
func GetBody(message *mail.Message) ([]byte, error) {
	body, err := ioutil.ReadAll(message.Body)
	if err != nil {
		log.Println("Error Reading message %s body: %s", message, err)
	}
	log.Printf("BODY: %s", body)
	return body, err

}

//PrintMessage prints te message to the log stream
func PrintMessage(messageID uint32, c *client.Client) {

	m, _ := GetMessage(messageID, "INBOX", c)

	header := m.Header
	for k, v := range header {
		log.Printf("Header[%s]: %s", k, v)
	}
	log.Println("Date:", header.Get("Date"))
	log.Println("From:", header.Get("From"))
	log.Println("To:", header.Get("To"))
	log.Println("Subject:", header.Get("Subject"))
}
