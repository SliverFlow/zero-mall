package logic

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/status"
	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"
	"server/app/user/model"
	"server/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.Nil, error) {
	var u model.UserGorm
	u.Password = common.BcryptHash(in.Password)
	u.Avatar = in.Avatar
	u.Email = in.Email
	u.UUID = uuid.NewV4().String()
	u.Username = in.Username
	u.Nickname = in.Nickname
	u.Role = in.Role

	err := l.svcCtx.UserModelGorm.Create(&u)
	if err != nil {
		return nil, status.Errorf(100001, "创建用户失败")
	}
	return &pb.Nil{}, nil
}
