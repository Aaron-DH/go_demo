package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	user := "dinghh@awcloud.com"
	password := "90op()OP"
	host := "smtp.exmail.qq.com:25"
	to := "344677472@.com"

	subject := "使用Golang发送邮件"

	body := `
        <html>
        <body>
        <h3>
        "Test send to email"
        </h3>
        </body>
        </html>
        `
	fmt.Println("Begin send email")
	err := sendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("->	Send mail error!", err)
	} else {
		fmt.Println("->	Send mail success!")
	}

}
