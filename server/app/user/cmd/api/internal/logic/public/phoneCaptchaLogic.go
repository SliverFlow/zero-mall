package public

import (
	"context"
	"server/common/aliyun"
	"server/common/xerr"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPhoneCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneCaptchaLogic {
	return &PhoneCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// PhoneCaptcha
// Author [SliverFlow]
// @desc 用户端发送短信验证码
func (l *PhoneCaptchaLogic) PhoneCaptcha(req *types.PhoneCaptcheReq) (resp *types.Nil, err error) {
	c := l.svcCtx.Config.AliSms
	err = aliyun.SendMsg(&aliyun.SmsConfig{
		SingName:        c.SingName,
		TemplateCode:    c.TemplateCode,
		AccessKeyID:     c.AccessKeyID,
		AccessKeySecret: c.AccessKeySecret,
		RedisCli:        l.svcCtx.Rsc,
		BaseStore:       c.BaseStore,
		BaseExpire:      c.BaseExpire,
	}, req.Phone)
	if err != nil {
		return nil, xerr.NewErrMsg("验证码发送失败")
	}

	return
}
