package base

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login
// Author [SliverFlow]
// @desc 系统登录
// @param req *types.LoginReq
// @return resp *types.LoginReply, err error
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	rep, err := l.svcCtx.UserRpc.Login(l.ctx, &userpb.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	u := rep.GetUser()
	if u.Role == 1 && u.Username != "admin" { // 普通用户
		return nil, xerr.NewErrMsg("此用户不能登录")
	}
	token, err := l.nextToken(u.UUID)
	if err != nil {
		return nil, xerr.NewErrMsg("token 获取失败")
	}

	var user types.User
	_ = copier.Copy(&user, u)

	return &types.LoginReply{
		Token: token,
		User:  user,
	}, nil
}

func (l *LoginLogic) nextToken(uuid string) (string, error) {
	conf := l.svcCtx.Config.XJwt
	jwt := xjwt.NewXJwt([]byte(conf.Secret), conf.Expire, conf.Buffer, conf.Isuser, conf.BlackListPrefix)
	token, err := jwt.SendToken(uuid)
	if err != nil {
		logx.Errorf("token create err", err.Error())
		return "", errors.New("token 创建失败")
	}
	return token, nil
}
