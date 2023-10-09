package public

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterByPhoneLogic {
	return &UserRegisterByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserRegisterByPhone
// Author [SliverFlow]
// @desc 用户端通过手机号注册
func (l *UserRegisterByPhoneLogic) UserRegisterByPhone(req *types.RegisterByPhoneReq) (resp *types.Nil, err error) {
	ok, err := l.checkCaptcha(req.Phone, req.Captcha)
	if err != nil || !ok {
		return nil, err
	}

	_, err = l.svcCtx.UserRpc.UserCheckPhone(l.ctx, &pb.PhoneReq{Phone: req.Phone})
	if err != nil {
		return nil, err
	}

	// 用户注册
	_, err = l.svcCtx.UserRpc.UserCreate(l.ctx, &pb.UserCreateReq{
		Role:  1,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}

	return
}

// checkCaptcha 验证验证码
func (l *UserRegisterByPhoneLogic) checkCaptcha(phone string, captcha string) (ok bool, err error) {
	c := l.svcCtx.Config.AliSms
	store := fmt.Sprintf("%s:%s", c.BaseStore, phone)
	result, err := l.svcCtx.Rsc.Get(l.ctx, store).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, xerr.NewErrMsg("验证码已经过期")
		}
		return false, err
	}
	if result != captcha {
		return false, xerr.NewErrMsg("验证码错误")
	}
	return true, nil
}
