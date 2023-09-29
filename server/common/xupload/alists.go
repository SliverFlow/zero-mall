package xupload

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/zeromicro/go-zero/core/logx"
)

// StsConfig
// Author [SliverFlow]
// @desc 阿里云 sts token 文件上传配置
type StsConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Duration        int64 // 临时 token 生效时间，单位秒
	RoleArn         string
	SessionName     string

	// 前端需要用的
	Region string
	Bucket string
}

func AliStsResponse(c *StsConfig) (*sts.AssumeRoleResponse, error) {

	client, _ := sts.NewClient(&openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &c.AccessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &c.AccessKeySecret,
		Endpoint:        &c.Endpoint,
	})

	response, err := client.AssumeRole(&sts.AssumeRoleRequest{
		DurationSeconds: &c.Duration,
		RoleArn:         &c.RoleArn,
		RoleSessionName: &c.SessionName,
	})
	if err != nil {
		logx.Error("aliyun oss sts  token error", err.Error())
		return nil, err
	}
	return response, err
}
