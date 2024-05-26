package email

import (
	"finpro-fenlie/config"

	"github.com/rs/zerolog/log"
	"gopkg.in/gomail.v2"
)

func Send(to, subject, body string) error {
	configData, err := config.InitEnv()
	if err != nil {
		log.Error().Err(err).Msg("failed to load config .env")
	}
	apiKey := configData.EmailSecret

	from := "752c23001@smtp-brevo.com"
	host := "smtp-relay.brevo.com"
	port := 587

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", "<h1>Account</h1><pre>"+string(body)+"</pre>")

	n := gomail.NewDialer(host, port, from, apiKey)

	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
