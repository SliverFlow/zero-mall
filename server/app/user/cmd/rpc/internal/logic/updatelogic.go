package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
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
	var u model.User
	_ = copier.Copy(&u, in)
	u.ID = uint(int(in.ID))

	err := l.svcCtx.UserModel.Update(l.ctx, &u)
	// 返回nil表示成功
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(100001, "user not exist in userId = %s", in.GetID())
		}
		return nil, err
	}
	return &pb.Nil{}, nil
}
