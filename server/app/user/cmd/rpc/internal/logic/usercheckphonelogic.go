package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCheckPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCheckPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCheckPhoneLogic {
	return &UserCheckPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserCheckPhone
// Author [SliverFlow]
// @desc 验证手机号是否注册
func (l *UserCheckPhoneLogic) UserCheckPhone(in *pb.PhoneReq) (*pb.UserNil, error) {
	ok, err := l.svcCtx.UserModel.CheckPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, status.Errorf(100001, "手机号验证失败")
	}
	if !ok {
		return nil, status.Errorf(100001, "此手机号已经注册注册，请使用验证码或者密码登录")
	}

	return &pb.UserNil{}, nil
}
