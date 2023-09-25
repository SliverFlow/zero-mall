package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"server/app/user/cmd/rpc/internal/svc"
	"server/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BusinessFindByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessFindByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessFindByUUIDLogic {
	return &BusinessFindByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BusinessFindByUUID
// Author [SliverFlow]
// @desc 通过 uuid 查询商户信息
func (l *BusinessFindByUUIDLogic) BusinessFindByUUID(in *pb.BusinessUUIDReq) (*pb.Business, error) {
	enter, err := l.svcCtx.UserModel.BusinessFindByUUID(l.ctx, in.GetUUID())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logx.Errorf("business not exist in UUID = %s", in.GetUUID())
			return nil, status.Errorf(100001, "不存在此商户")
		}
		return nil, err
	}

	var pbbusiness pb.Business
	_ = copier.Copy(&pbbusiness, enter)
	pbbusiness.CreatedAt = enter.CreatedAt.Unix()
	pbbusiness.UpdatedAt = enter.UpdatedAt.Unix()

	return &pbbusiness, nil
}
