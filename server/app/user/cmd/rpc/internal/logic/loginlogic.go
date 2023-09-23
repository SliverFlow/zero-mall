package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"server/common"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login
// Author [SliverFlow]
// @desc 用户登录
func (l *LoginLogic) Login(in *pb.UserLoginReq) (*pb.UserLoginReply, error) {

	user, err := l.svcCtx.UserModel.UserFindByUsername(l.ctx, in.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "not find user by username is %s", in.Username)
		}
		return nil, err
	}
	if user.Status == 0 {
		return nil, status.Errorf(100001, "用户已经被禁用，请联系管理员解除禁用")
	}
	if ok := common.BcryptCheck(in.Password, user.Password); !ok { // 密码比对未成功
		return nil, status.Errorf(100001, "password check not success")
	}
	u := copyUserFoReply(user) // 生成返回结构
	return &pb.UserLoginReply{
		User: u,
	}, nil
}
