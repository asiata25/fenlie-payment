package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func Send(to, subject, body string) error {
	host := "smtp.gmail.com"
	port := 587
	sender := "andreekapradana@gmail.com"

	// FIXME: cann't get email secret from env
	password := "ovnr bbnh tvdb ssyp"

	msg := gomail.NewMessage()
	msg.SetHeader("From", fmt.Sprintf("%s <%s>", "Fenlie", sender))
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, sender, password)

	if err := d.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
