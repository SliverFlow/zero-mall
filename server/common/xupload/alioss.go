package xupload

import (
	"fmt"
	"mime/multipart"
)

type Alioss struct {
	endpoint        string
	accessKeyId     string
	accessKeySecret string
	bucketName      string
}

func (a *Alioss) UploadFile(file *multipart.FileHeader) (string, string, error) {
	fmt.Println("gnvfioqerbngvire", a)
	return "", "", nil
}

func newAlioss(c *AliossConfig) *Alioss {
	return &Alioss{
		endpoint:        c.Endpoint,
		accessKeyId:     c.AccessKeyId,
		accessKeySecret: c.AccessKeySecret,
		bucketName:      c.Endpoint,
	}
}
