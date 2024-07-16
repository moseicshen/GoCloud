package helper

import (
	"GoCloud/core/define"
	"GoCloud/core/key"
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid2 "github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/tencentyun/cos-go-sdk-v5"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"time"
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

func RandCode() string {
	s := "0123456789"
	code := ""
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < define.EmailCodeLength; i++ {
		code += string(s[r.Intn(len(s))])
	}
	return code
}

func UUID() string {
	uuid, err := uuid2.NewUUID()
	if err != nil {
		return ""
	}
	return uuid.String()
}

// UploadFile upload to Tencent Cloud COS
func UploadFile(r *http.Request) (string, error) {
	u, _ := url.Parse(key.COSURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  key.SecretID,
			SecretKey: key.SecretKey,
		},
	})
	file, fileHeader, err := r.FormFile("file")
	uuid := UUID()
	name := "GoCloud/" + uuid + path.Ext(fileHeader.Filename)
	_, err = c.Object.Put(context.Background(), name, file, nil)
	if err != nil {
		panic(err)
	}
	keyStr := key.COSURL + "/" + name
	return keyStr, nil
}
