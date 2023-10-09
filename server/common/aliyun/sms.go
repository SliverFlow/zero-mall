package aliyun

import (
	"context"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"strconv"
	"time"
)

// 阿里云短信服务

func createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, err = dysmsapi20170525.NewClient(config)
	return _result, err
}

// GetAuthCode 获取验证码
func getAuthCode() string {
	// 生成 6 位随机数
	randomNum := rand.Intn(900000) + 100000
	return strconv.Itoa(randomNum)
}

// SendMsg 发送短信
func SendMsg(c *SmsConfig, phone string) error {
	// 请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID 和 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	client, err := createClient(tea.String(c.AccessKeyID), tea.String(c.AccessKeySecret))
	if err != nil {
		return err
	}

	code := getAuthCode()

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(c.SingName),
		TemplateCode:  tea.String(c.TemplateCode),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	resp, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		return err
	}

	// 将验证码存在redis中
	if err := setAuthCode(c.RedisCli, c.BaseStore, phone, code, c.BaseExpire); err != nil {
		logx.Error("set captcha for redis err", err.Error())
		return err
	}
	logx.Info(util.ToJSONString(resp))
	return err
}

type SmsConfig struct {
	SingName        string
	TemplateCode    string
	AccessKeyID     string
	AccessKeySecret string
	RedisCli        *redis.Client
	BaseStore       string
	BaseExpire      int
}

// 将验证码存在redis中
func setAuthCode(cli *redis.Client, baseStore string, phone string, code string, exp int) error {
	store := fmt.Sprintf("%s:%s", baseStore, phone)
	return cli.Set(context.Background(), store, code, time.Duration(exp)*time.Second).Err()
}
