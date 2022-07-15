package main

import (
	"github.com/wandachu/bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"time"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MaliChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "127.0.0.1"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, "Hello, <strong>world</strong>!")

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("email sent")
	}
}
