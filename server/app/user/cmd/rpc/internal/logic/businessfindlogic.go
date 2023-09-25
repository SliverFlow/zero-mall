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

type BusinessFindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBusinessFindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BusinessFindLogic {
	return &BusinessFindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BusinessFind
// Author [SliverFlow]
// @desc 通过 商户id 查询商户信息
func (l *BusinessFindLogic) BusinessFind(in *pb.BusinessIDReq) (*pb.Business, error) {
	enter, err := l.svcCtx.UserModel.BusinessFind(l.ctx, in.GetBusinessID())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logx.Errorf("business not exist in BusinessID = %s", in.GetBusinessID())
			return nil, status.Errorf(100001, "不存在此商家")
		}
		return nil, err
	}

	var pbbusiness pb.Business
	_ = copier.Copy(&pbbusiness, enter)
	pbbusiness.CreatedAt = enter.CreatedAt.Unix()
	pbbusiness.UpdatedAt = enter.UpdatedAt.Unix()

	return &pbbusiness, nil
}
