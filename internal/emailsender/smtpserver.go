package emailsender

import (
	"log"
	"net/mail"
	"net/smtp"
)

type smtpServer interface {
	send(*email)
}

type serverDetails struct {
	auth smtp.Auth
	URL  string
}

func (sd *serverDetails) send(eml *email) {
	for _, e := range eml.to {
		if !validEmail(e) {
			log.Printf("%s is not a valid email address", e)
			return
		}
	}

	err := smtp.SendMail(sd.URL, sd.auth, eml.from, eml.to, eml.message)
	if err != nil {
		log.Println(err)
	}
}

func ConfigureServer(cfg *smtpConfig) smtpServer {
	return &serverDetails{
		auth: smtp.PlainAuth("", cfg.username, cfg.password, cfg.smtpHost),
		URL:  cfg.smtpHost + ":" + cfg.smtpPort,
	}
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
