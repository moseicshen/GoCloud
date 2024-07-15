package helper

import (
	"GoCloud/core/define"
	"GoCloud/core/key"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity string, name string) (tokenStr string, err error) {
	userClaim := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaim)
	tokenStr, err = tk.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return
}

func SendEmailCode(userEmail string, code string) (err error) {
	sender := fmt.Sprintf("GoCloud-No-Reply<%s>", key.EmailSender)
	htmlStr := fmt.Sprintf("<h1>Your Test Code is:</h1><h2>%s</h2>", code)
	addrStr := fmt.Sprintf("%s%s", key.EmailServer, key.EmailPort)
	e := email.NewEmail()
	e.From = sender
	e.To = []string{userEmail}
	e.Subject = "Go Cloud Test Email"
	e.HTML = []byte(htmlStr)
	err = e.SendWithTLS(addrStr, smtp.PlainAuth("", key.EmailSender, key.EmailKey, key.EmailServer),
		&tls.Config{InsecureSkipVerify: true, ServerName: key.EmailServer})
	if err != nil {
		return err
	}
	return
}
