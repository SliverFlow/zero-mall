package public

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
	"server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginByPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginByPhoneLogic {
	return &UserLoginByPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserLoginByPhone
// Author [SliverFlow]
// @desc 用户端使用电话号码加验证码登录
func (l *UserLoginByPhoneLogic) UserLoginByPhone(req *types.UserLoginByPhoneReq) (resp *types.UserLoginReply, err error) {

	ok, err := l.checkCaptcha(req.Phone, req.Captcha)
	if err != nil || !ok {
		return nil, err
	}

	pbreply, err := l.svcCtx.UserRpc.UserFindByPhone(l.ctx, &pb.PhoneReq{Phone: req.Phone})
	if err != nil {
		return nil, err
	}

	user := pbreply.GetUser()

	if user.GetStatus() == 0 {
		return nil, xerr.NewErrMsg("当前用户已被禁用，请联系管理员进行解禁")
	}
	if user.Role != 1 {
		return nil, xerr.NewErrMsg("当前账号不能登录用户端，请注册用户端账号")
	}

	// 派发 token
	token, err := l.nextToken(user.GetUUID())
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}

	return &types.UserLoginReply{
		Token: token,
		User: types.UserReply{
			UUID:     user.GetUUID(),
			Username: user.GetUsername(),
			Nickname: user.GetNickname(),
			Avatar:   user.GetAvatar(),
			Phone:    user.GetPhone(),
			Email:    user.GetEmail(),
		},
	}, nil
}

func (l *UserLoginByPhoneLogic) nextToken(uuid string) (string, error) {
	conf := l.svcCtx.Config.XJwt
	jwt := xjwt.NewXJwt([]byte(conf.Secret), conf.Expire, conf.Buffer, conf.Isuser, conf.BlackListPrefix)
	token, err := jwt.SendToken(uuid)
	if err != nil {
		logx.Errorf("token create err", err.Error())
		return "", errors.New("token 创建失败")
	}
	return token, nil
}

// checkCaptcha 验证验证码
func (l *UserLoginByPhoneLogic) checkCaptcha(phone string, captcha string) (ok bool, err error) {
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
