package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	ssh.Handle(func(sess ssh.Session) {

		io.WriteString(sess, fmt.Sprintf("Hello %s\n", sess.User()))

		term := terminal.NewTerminal(sess, "> ")
		for {
			line, err := term.ReadLine()
			log.Println(line)
			if err != nil {
				break
			}
		}
		log.Println("terminal closed")
	})
	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
