package public

import (
	"context"
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
	conf := l.svcCtx.Config.XJwt
	// todo: add your logic here and delete this line
	jwt := xjwt.NewXJwt([]byte(conf.Secret), conf.Expire, conf.Buffer, conf.Isuser, conf.BlackListPrefix)
	token, err := jwt.SendToken("172781", "bfuewf")
	if err != nil {
		return nil, xerr.NewErrMsg(err.Error())
	}
	return &types.LoginReply{Token: token}, nil
}
