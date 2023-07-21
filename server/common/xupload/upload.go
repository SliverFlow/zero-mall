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
	// UploadFile 文件 保存的位置 返回 文件名 文件访问 utl
	UploadFile(file *multipart.FileHeader) (name string, url string, err error)
	// UploadFileByStorePath(file *multipart.FileHeader, store string) (name string, url string, err error)
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
