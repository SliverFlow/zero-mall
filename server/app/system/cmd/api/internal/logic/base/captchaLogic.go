package base

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"
	"server/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var store = base64Captcha.DefaultMemStore

func (l *CaptchaLogic) Captcha(req *types.NilReq) (resp *types.CaptchaReply, err error) {
	c := l.svcCtx.Config.Captcha
	driver := base64Captcha.NewDriverDigit(c.ImgHeight, c.ImgWidth, c.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, base64, err := cp.Generate()
	if err != nil {
		logx.Error("captcha generate err:", err.Error())
		return nil, xerr.NewErrMsg("captcha generate err")
	}
	return &types.CaptchaReply{
		CaptchaId:     id,
		PicPath:       base64,
		CaptchaLength: c.KeyLong,
	}, nil
}
