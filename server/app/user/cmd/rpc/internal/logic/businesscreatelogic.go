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

type BusinessCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessCreateLogic {
	return &BusinessCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BusinessCreate
// Author [SliverFlow]
// @desc 创建商家
func (l *BusinessCreateLogic) BusinessCreate(in *pb.BusinessCreateReq) (*pb.UserNil, error) {
	// 验证用户信息
	user, err := l.svcCtx.UserModel.UserFindByUUID(l.ctx, in.GetUUID())
	if err != nil {
		logx.Error("find business admin user err", err.Error())
		return nil, status.Errorf(100001, "此用户不存在")
	}
	// 验证用户是否已经是商家
	userList, _ := l.svcCtx.UserModel.UserListFindByUsername(l.ctx, user.Username)
	if len(userList) > 0 {
		for i := range userList {
			if userList[i].Role == 2 {
				return nil, status.Errorf(100001, "此用户已经是商家")
			}
		}
	}
	// 生成唯一 id
	uid := uuid.NewV4().String()
	// 生成用户信息
	newUser := model.User{UUID: uid, Username: user.Username, Email: user.Email, Nickname: user.Nickname, Password: user.Password, Avatar: user.Avatar, Role: 2, Status: 0, Phone: user.Phone}
	// 创建信息
	if err = l.svcCtx.UserModel.UserCreate(l.ctx, &newUser); err != nil {
		return nil, err
	}
	var bu model.Business
	bu.UUID = uid
	bu.BusinessID = strconv.FormatInt(common.SnowWorker.NextId(), 10)
	bu.Name = "默认店铺名称"
	bu.Detail = "默认店铺描述"
	// 创建商家用户
	if err = l.svcCtx.UserModel.BusinessCreate(l.ctx, &bu); err != nil {
		logx.Error("create business err", err.Error())
		return nil, status.Errorf(100001, "创建商家失败")
	}
	return &pb.UserNil{}, nil
}
