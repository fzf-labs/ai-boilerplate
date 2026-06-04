package service

import (
	"crypto/tls"
	"fmt"
	"mime"
	"net/smtp"
	"strconv"
	"strings"

	"github.com/fzf-labs/ai-boilerplate-backend/internal/data/gorm/ai_boilerplate_model"
	"github.com/fzf-labs/goutil/uuidutil"
)

func buildSMTPMessage(fromMail, fromName, toMail, subject, body string) []byte {
	fromHeader := fromMail
	if strings.TrimSpace(fromName) != "" {
		fromHeader = fmt.Sprintf("%s <%s>", mime.QEncoding.Encode("utf-8", fromName), fromMail)
	}
	subjectHeader := mime.QEncoding.Encode("utf-8", subject)
	message := strings.Join([]string{
		"From: " + fromHeader,
		"To: " + toMail,
		"Subject: " + subjectHeader,
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=UTF-8",
		"",
		body,
	}, "\r\n")
	return []byte(message)
}

func sendSMTPMail(account *ai_boilerplate_model.MailAccount, fromName, toMail, subject, body string) (string, error) {
	messageID := uuidutil.GenUUID()
	if account == nil {
		return messageID, fmt.Errorf("mail account is nil")
	}

	addr := account.Host + ":" + strconv.Itoa(int(account.Port))
	message := buildSMTPMessage(account.Mail, fromName, toMail, subject, body)
	var auth smtp.Auth
	if strings.TrimSpace(account.Username) != "" || strings.TrimSpace(account.Password) != "" {
		auth = smtp.PlainAuth("", account.Username, account.Password, account.Host)
	}

	if !account.SslEnable {
		if err := smtp.SendMail(addr, auth, account.Mail, []string{toMail}, message); err != nil {
			return messageID, err
		}
		return messageID, nil
	}

	conn, err := tls.Dial("tcp", addr, &tls.Config{ServerName: account.Host})
	if err != nil {
		return messageID, err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, account.Host)
	if err != nil {
		return messageID, err
	}
	defer client.Close()

	if auth != nil {
		if err := client.Auth(auth); err != nil {
			return messageID, err
		}
	}
	if err := client.Mail(account.Mail); err != nil {
		return messageID, err
	}
	if err := client.Rcpt(toMail); err != nil {
		return messageID, err
	}
	writer, err := client.Data()
	if err != nil {
		return messageID, err
	}
	if _, err := writer.Write(message); err != nil {
		_ = writer.Close()
		return messageID, err
	}
	if err := writer.Close(); err != nil {
		return messageID, err
	}
	if err := client.Quit(); err != nil {
		return messageID, err
	}
	return messageID, nil
}
