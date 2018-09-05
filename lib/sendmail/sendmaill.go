package sendmail

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

func encodeRFC2047(str string) string {
	addr := mail.Address{Name: str, Address: ""}
	return strings.Trim(addr.String(), " <>")
}

/**
 * 邮件发送
 * @param toMail  		邮件接收者账号
 * @param Subject		邮件主题
 * @param fromName		邮件发送着名称（便于接收者直观查看邮件来源)
 * @param mailtype      邮件内容类型： html | plain
 * @param body			邮件内容
 *
 *
 */
func SendMail(toMail, Subject, fromName, mailtype, body string) error {

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		"starrycustomerservice@gmail.com",
		"starry2018",
		"smtp.gmail.com",
	)
	from := mail.Address{Name: fromName, Address: "starrycustomerservice@gmail.com"}
	to := mail.Address{Name: "收件人", Address: toMail}

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(Subject)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/" + mailtype + "; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	fmt.Println("start send mail to : " + toMail)
	err := smtp.SendMail(
		smtpServer+":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)
	fmt.Println("finished sen mail.")

	return err
}
