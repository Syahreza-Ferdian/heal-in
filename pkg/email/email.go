package email

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Syahreza-Ferdian/heal-in/model"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailService interface {
	SendEmail(user *model.UserRegister, data *EmailData)
}

type EmailSender struct {
	Name      string
	FromEmail string
	Password  string
}

type EmailData struct {
	RedirectURL string
	FirstName   string
	Subject     string
	WebURL      string
}

func NewEmailSender(name, fromEmail, password string) EmailService {
	return &EmailSender{
		Name:      name,
		FromEmail: fromEmail,
		Password:  password,
	}
}

func (e *EmailSender) SendEmail(user *model.UserRegister, data *EmailData) {
	serverPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	// parse template
	tmpl, err := ParseTemplateDir("../../pkg/email/template")
	if err != nil {
		panic(err)
	}

	// create a new buffer
	var body bytes.Buffer

	// execute the template and write it to the buffer
	tmpl.ExecuteTemplate(&body, "base.html", &data)

	// send email
	// err = smtp.SendMail(ServerAddress,
	// 	smtp.PlainAuth("", e.Name, e.Password, AuthAddress),
	// 	e.FromEmail, []string{user.Email}, []byte(body.String()))

	m := gomail.NewMessage()
	m.SetHeader("From", e.FromEmail)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	dialer := gomail.NewDialer(os.Getenv("SMTP_HOST"), serverPort, e.Name, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = dialer.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
