package email

import (
	"bytes"
	"log"
	"path"
	"text/template"

	"github.com/MociW/store-api-golang/pkg/config"
	"gopkg.in/gomail.v2"
)

type emailService struct {
	cnf *config.Config
}

func NewEmailService(cnf *config.Config) EmailService {
	return &emailService{cnf: cnf}
}

func (e *emailService) Send(to string, subject string, body EmailData) error {
	var filepath = path.Join("template", "otp.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Error parsing template: %s", err.Error())
		return err
	}

	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, body); err != nil {
		log.Printf("Error executing template: %s", err.Error())
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", e.cnf.Mail.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", bodyBuffer.String())

	// Set up the SMTP server configuration
	d := gomail.NewDialer(e.cnf.Mail.Host, e.cnf.Mail.Port, e.cnf.Mail.User, e.cnf.Mail.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email: %s", err.Error())
		return err
	}
	return nil
}
