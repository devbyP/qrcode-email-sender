package main

import (
	"log"

	"github.com/devbyP/qrcode-email-sender/pkgs/emailSender"
	"github.com/devbyP/qrcode-email-sender/pkgs/qrcodeGen"
)

func main() {
	qrfile, err := qrcodeGen.GenerateCode("never", "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	if err != nil {
		log.Fatal(err)
	}
	if err = emailSender.SendEmail("bytedopdap@gmail.com", qrfile); err != nil {
		log.Fatal(err)
	}
}
