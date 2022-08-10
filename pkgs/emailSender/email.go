package emailSender

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, qrcode string) error {
	sender := os.Getenv("EMAIL_APP_USER")
/* 	var port int
	var err error */
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "test email title")
	
	// smtp.gmail.com -> google's smtp server
/* 	if port, err = strconv.Atoi(os.Getenv("SMTP_PORT")); err != nil {
		return err
	} */
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return fmt.Errorf("SendEmail: %w", err)
	}
	dialer := gomail.NewDialer("smtp.gmail.com", port, sender, os.Getenv("EMAIL_APP_PASSWORD"))
	return dialer.DialAndSend(m)
}

func fileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type emailMessage struct {
	mess *gomail.Message
}

func (em *emailMessage) Write(b []byte) (int, error) {
	em.mess.SetBody("text/html", string(b))
	return len(b), nil
}
