package public

import (
	"context"
	"github.com/pkg/errors"
	"server/app/user/cmd/api/internal/config"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {

	logx.Info("Login with user ", req.Username)
	logx.Error("Login with user ", req.Password)
	token, err := l.nextToken(l.svcCtx.Config)
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	return &types.LoginReply{Token: token}, nil
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
