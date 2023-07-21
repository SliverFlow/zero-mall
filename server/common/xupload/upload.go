package xupload

import "mime/multipart"

type Config struct {
	Type  string
	Local struct {
		Path string
	}
	Alioss *AliossConfig
}

type AliossConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

type XUpload interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	//DeleteFile(key string) error
}

func NewOss(c *Config) XUpload {
	switch c.Type {
	case "alioss":
		return newAlioss(c.Alioss)
	default:
		return newAlioss(c.Alioss)
	}

}
