package logic

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/status"
	"server/app/user/model"
	"server/common"
	"strconv"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

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

var worker = common.NewWorker(1)

func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.Nil, error) {
	var u model.User
	u.Password = common.BcryptHash(in.Password)
	u.Avatar = in.Avatar
	u.Email = in.Email
	u.Uuid = uuid.NewV4().String()
	u.UserId = strconv.FormatInt(worker.NextId(), 10)
	u.Username = in.Username
	u.Nickname = in.Nickname

	err := l.svcCtx.UserModel.Insert(l.ctx, nil, &u)
	if err != nil {
		return nil, status.Errorf(100001, "创建用户失败")
	}
	return &pb.Nil{}, nil
}
