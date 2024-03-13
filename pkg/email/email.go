package email

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendEmail(user *model.UserRegister, data *EmailData) error
}

type EmailSender struct {
	Name      string
	Password  string
}

type EmailData struct {
	RedirectURL string
	FirstName   string
	Subject     string
	WebURL      string
}

func NewEmailSender(name, password string) EmailService {
	return &EmailSender{
		Name:      name,
		Password:  password,
	}
}

func (e *EmailSender) SendEmail(user *model.UserRegister, data *EmailData) error {
	serverPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	body := fmt.Sprintf("<p>Hi, %s</p> <p>Follow this link below to confirm your email address. If you didn't create an account with <a href='%s'>Heal.in</a>, you can safely delete this email.</p> <a href='%s'>Klik</a> <br> <p>You received this email because we received a request for registration for your account. If you didn't request that registration into our service, you can safely delete this email.</p>", data.FirstName, data.WebURL, data.RedirectURL)

	m := gomail.NewMessage()
	m.SetHeader("From", e.Name)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body)
	m.AddAlternative("text/plain", html2text.HTML2Text(body))

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), serverPort, e.Name, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
