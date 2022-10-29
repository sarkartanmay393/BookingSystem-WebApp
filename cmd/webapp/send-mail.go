package main

import (
	"time"

	"github.com/sarkartanmay393/RoomReservation-WebApp/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenAndSendMails() {
	go func() {
		for {
			_ = <-app.MailChannel
			app.InfoLog.Println("email not sending in dev mode.")
			// sendEmail(email)
		}
	}()
}
func sendEmail(email models.MailData) {
	// Setting up our mail server with necessaries
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		app.ErrorLog.Println("Couldn't connect to SMTP server!\t", err)
		return
	}

	msg := mail.NewMSG()
	msg.SetFrom(email.From).AddTo(email.To).SetSubject(email.Subject)
	msg.SetBody(mail.TextHTML, email.Content)

	err = msg.Send(client)
	if err != nil {
		app.ErrorLog.Println("Unable to send email! \t", err)
		return
	}
	app.InfoLog.Println("Mail sent successfully!")
}
