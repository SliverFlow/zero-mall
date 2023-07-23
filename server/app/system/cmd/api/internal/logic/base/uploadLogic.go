package base

import (
	"context"
	"net/http"
	"server/app/system/cmd/api/internal/config"
	"server/common/xerr"
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

const maxFileSize = 10 << 20 // 10 MB
func (l *UploadLogic) Upload(r *http.Request) (resp *types.FileUploadReply, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	_, file, err := r.FormFile("file")
	if err != nil {
		logx.Error("file get err", err)
		return nil, err
	}
	conf := l.getOssConf(l.svcCtx.Config)
	oss := xupload.NewOss(conf)
	name, url, err := oss.UploadFile(file)
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	return &types.FileUploadReply{
		FileStoreName: name,
		Name:          file.Filename,
		Url:           url,
		FileSize:      file.Size,
	}, nil
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
