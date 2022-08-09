package emailSender

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, qrcode string) error {
/* 	var port int
	var err error */
	m := gomail.NewMessage()
	m.SetHeader("From", "thanaponpuipui@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "test email title")
	m.Embed("./codeimg/"+qrcode)
	m.SetBody("text/html", `
		<h1>test header</h1>
		<p>test body paragraph</p>
		<img src="cid:`+qrcode+`" alt="qrcode image" />
	`)
	// smtp.gmail.com -> google's smtp server
/* 	if port, err = strconv.Atoi(os.Getenv("SMTP_PORT")); err != nil {
		return err
	} */
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "thanaponpuipui@gmail.com", "erpqdhzlvnwkpgwf")
	return dialer.DialAndSend(m)
}

func fileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
