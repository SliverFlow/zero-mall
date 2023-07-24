package base

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"server/app/system/cmd/api/internal/config"
	"server/app/system/cmd/rpc/pb"
	"server/common/xerr"
	"server/common/xjwt"
	"time"

	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"

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
	rep, err := l.svcCtx.SystemRpc.Login(l.ctx, &pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	token, err := l.nextToken(l.svcCtx.Config)
	if err != nil {
		return nil, xerr.NewErrMsg("token 获取失败")
	}
	u := rep.GetUser()
	fmt.Println("time", time.Unix(u.UpdatedAt, 0))
	return &types.LoginReply{
		Token: token,
		User: types.User{
			Username:  u.Username,
			UUID:      u.UUID,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Avatar:    u.Avatar,
			CreatedAt: u.CreatedAt,
			UpdetedAt: u.UpdatedAt,
		},
	}, nil
}

func (l *LoginLogic) nextToken(c config.Config) (string, error) {
	conf := l.svcCtx.Config.XJwt
	jwt := xjwt.NewXJwt([]byte(conf.Secret), conf.Expire, conf.Buffer, conf.Isuser, conf.BlackListPrefix)
	token, err := jwt.SendToken("172781", "bfuewf")
	if err != nil {
		logx.Errorf("token create err", err.Error())
		return "", errors.New("token 创建失败")
	}
	return token, nil
}
