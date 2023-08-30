package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"server/app/user/model"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Update 更新用户信息
func (l *UpdateLogic) Update(in *pb.UpdateReq) (*pb.Nil, error) {
	err := l.svcCtx.UserModel.Update(l.ctx, nil, &model.User{
		Username: in.Username,
		Email:    in.Email,
		Nickname: in.Nickname,
		Password: in.Password,
		Avatar:   in.Avatar,
		Status:   in.Status,
	})
	// 返回nil表示成功
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "user not exist in userId = %s", in.GetID())
		}
		return nil, err
	}
	return &pb.Nil{}, nil
}
