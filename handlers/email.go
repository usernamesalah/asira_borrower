package handlers

import (
	"asira_borrower/asira"
	"fmt"

	"gopkg.in/gomail.v2"
)

func sendMail(to string, subject, message string) error {
	Config := asira.App.Config.GetStringMap(fmt.Sprintf("%s.mailer", asira.App.ENV))
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", Config["email"].(string))
	mailer.SetHeader("To", to)
	mailer.SetAddressHeader("Cc", "no-reply@ayannah.com", "git Ayannah Support")
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewPlainDialer(Config["host"].(string),
		Config["port"].(int),
		Config["email"].(string),
		Config["password"].(string))

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
