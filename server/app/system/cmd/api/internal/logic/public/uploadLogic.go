package public

import (
	"context"
	"fmt"
	"server/app/system/cmd/api/internal/config"
	"server/common/xupload"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload() (resp *types.FileUploadReply, err error) {
	// todo: add your logic here and delete this line
	conf := l.getOssConf(l.svcCtx.Config)
	fmt.Println(conf.Alioss.BucketName)
	oss := xupload.NewOss(conf)
	_, _, err = oss.UploadFile(nil)
	if err != nil {
		fmt.Println("gn9ueqrhgburegv", err)
	}
	fmt.Println(oss)
	return
}

// 获取 oss 配置
func (l *UploadLogic) getOssConf(conf config.Config) *xupload.Config {
	c := conf.Oss
	return &xupload.Config{
		Type: c.Type,
		Local: struct {
			Path string
		}{
			Path: c.Local.Path,
		},
		Alioss: &xupload.AliossConfig{
			Endpoint:        c.Alioss.Endpoint,
			AccessKeyId:     c.Alioss.AccessKeyId,
			AccessKeySecret: c.Alioss.AccessKeySecret,
			BucketName:      c.Alioss.BucketName,
		},
	}
}
