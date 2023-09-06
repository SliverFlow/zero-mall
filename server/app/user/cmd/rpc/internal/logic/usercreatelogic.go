package logic

import (
	"context"
	"github.com/jinzhu/copier"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/status"
	"server/app/user/model"
	"server/common"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *pb.CreateReq) (*pb.Nil, error) {
	var u model.User
	_ = copier.Copy(&u, in)
	u.Password = common.BcryptHash(in.Password)
	u.UUID = uuid.NewV4().String()

	err := l.svcCtx.UserModel.UserCreate(l.ctx, &u)
	if err != nil {
		return nil, status.Errorf(100001, "创建用户失败")
	}
	return &pb.Nil{}, nil
}
