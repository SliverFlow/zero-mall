package public

import (
	"context"
	"github.com/pkg/errors"
	"server/app/user/cmd/rpc/pb"
	"server/common"
	"server/common/xerr"
	"server/common/xjwt"

	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginByPhoneOrUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginByPhoneOrUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginByPhoneOrUsernameLogic {
	return &UserLoginByPhoneOrUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserLoginByPhoneOrUsername
// Author [SliverFlow]
// @desc 用户端使用用户名加密码获取电话号码加密码的形式登录
func (l *UserLoginByPhoneOrUsernameLogic) UserLoginByPhoneOrUsername(req *types.LoginByPhoneOrUsernameReq) (resp *types.UserLoginReply, err error) {
	pbreply, err := l.svcCtx.UserRpc.UserFindByPhoneOrUsername(l.ctx, &pb.UserFindByPhoneOrUsernameReq{Username: req.Username})
	if err != nil {
		return nil, err
	}

	user := pbreply.GetUser()
	// 比对密码
	if ok := common.BcryptCheck(req.Password, user.GetPassword()); !ok {
		return nil, xerr.NewErrMsg("密码输入有误，请重新输入")
	}

	if user.GetStatus() == 0 {
		return nil, xerr.NewErrMsg("当前用户已被禁用，请联系管理员进行解禁")
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

func (l *UserLoginByPhoneOrUsernameLogic) nextToken(uuid string) (string, error) {
	conf := l.svcCtx.Config.XJwt
	jwt := xjwt.NewXJwt([]byte(conf.Secret), conf.Expire, conf.Buffer, conf.Isuser, conf.BlackListPrefix)
	token, err := jwt.SendToken(uuid)
	if err != nil {
		logx.Errorf("token create err", err.Error())
		return "", errors.New("token 创建失败")
	}
	return token, nil
}
