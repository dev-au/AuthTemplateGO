package utils

import (
	"AuthTemplate/src"
	"fmt"
	"log"
	"net/smtp"
)

func sendEmail(receiverEmail []string, subject, template string) {

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := []byte(fmt.Sprintf("%s \r\n"+
		"\r\n"+
		"%s \r\n", subject, template))

	auth := smtp.PlainAuth("", src.Config.GmailAccount, src.Config.GmailPassword, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, src.Config.GmailAccount, receiverEmail, msg)
	if err != nil {
		log.Fatal("Failed to send email:", err)
	}
}

func SendVerifyLink(receiver string, data map[string]interface{}) {

	cipher, err := EncryptAES(data)
	if err != nil {
		log.Fatal("Failed to encrypt link:", err)
	}

	link := fmt.Sprintf("%suser/verify/%s", src.Config.ApplicationUrl, cipher)

	template := fmt.Sprintf("Your link is: " + link)

	sendEmail([]string{receiver}, "One Time Verifier Url", template)
}
