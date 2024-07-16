package test

import (
	"GoCloud/core/key"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestCos(t *testing.T) {
	u, _ := url.Parse(key.COSURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  key.SecretID,
			SecretKey: key.SecretKey,
		},
	})
	name := "GoCloud/example.png"
	// 2. 通过本地文件上传对象
	_, err := c.Object.PutFromFile(context.Background(), name, "../files/Factory.png", nil)
	if err != nil {
		panic(err)
	}
}

func TestCosUploadByReader(t *testing.T) {
	u, _ := url.Parse(key.COSURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  key.SecretID,
			SecretKey: key.SecretKey,
		},
	})
	name := "GoCloud/example2.png"
	// 3.通过文件流上传对象
	fd, err := os.Open("../files/Factory.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	_, err = c.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		panic(err)
	}

}
