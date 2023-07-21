package xupload

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
	"path"
)

type Alioss struct {
	endpoint        string
	accessKeyId     string
	accessKeySecret string
	bucketName      string
}

func newAlioss(c *AliossConfig) *Alioss {
	return &Alioss{
		endpoint:        c.Endpoint,
		accessKeyId:     c.AccessKeyId,
		accessKeySecret: c.AccessKeySecret,
		bucketName:      c.BucketName,
	}
}

// getBucket 获取数据桶对象
func (a *Alioss) getBucket() (*oss.Bucket, error) {
	// 获取连接
	client, cliErr := oss.New(a.endpoint, a.accessKeyId, a.accessKeySecret)
	if cliErr != nil {
		return nil, errors.New("function oss.New() filed err " + cliErr.Error())
	}
	// 判断存储桶是否存在
	exist, _ := client.IsBucketExist(a.bucketName)
	if !exist {
		// 存储桶不存在 创建存储桶
		buckErr := client.CreateBucket(a.bucketName)
		if buckErr != nil {
			return nil, errors.New("function client.CreateBucket() filed err :" + buckErr.Error())
		}
		// 设置桶的读写权限
		buckErr = client.SetBucketACL(a.bucketName, oss.ACLPublicReadWrite)
		if buckErr != nil {
			return nil, errors.New("function client.CreateBucket() filed err :" + buckErr.Error())
		}

	}
	// 获取桶对象
	bucket, buckErr := client.Bucket(a.bucketName)
	if buckErr != nil {
		return nil, errors.New("function client.Bucket() filed err :" + buckErr.Error())
	}

	return bucket, nil
}

func (a *Alioss) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := a.getBucket()
	if err != nil {
		logx.Error("获取桶对象失败", err)
		return "", "", err
	}
	// 获取文件后缀
	fileSuffix := path.Ext(file.Filename)
	// 获取新文件名
	fileNewName := uuid.NewV4().String() + fileSuffix
	// 将文件上传
	open, err := file.Open()
	if err != nil {
		logx.Error("文件打开失败", err)
		return "", "", err
	}
	defer open.Close()
	// 上传
	err = bucket.PutObject(fileNewName, open)
	if err != nil {
		logx.Error("文件上传至 aliyunoss 失败", err)
		return "", "", err
	}
	// 获取访问路径
	url := fmt.Sprintf("https://%s.%s/%s", a.bucketName, a.endpoint, fileNewName)
	return fileNewName, url, nil
}
