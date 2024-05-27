package email

import (
	"finpro-fenlie/config"
	"fmt"

	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
)

func Send(to, subject, body string) error {
	configData, err := config.InitEnv()
	if err != nil {
		log.Error().Err(err).Msg("failed to load config .env")
	}

	host := "smtp.gmail.com"
	port := 587
	sender := "andreekapradana@gmail.com"
	password := configData.EmailSecret

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
