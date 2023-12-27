package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"os"
	"server/app/order/cmd/rpc/internal/config"
	"server/app/order/model"
	pruductpb "server/app/product/cmd/rpc/pb"
	"server/app/product/cmd/rpc/product"
	userpb "server/app/user/cmd/rpc/pb"
	"server/app/user/cmd/rpc/user"
	"server/common"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    userpb.UserClient
	ProductRpc pruductpb.ProductClient

	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := common.InitDB(c.Mysql)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		os.Exit(0)
	}
	// autoMigrate(db)

	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(db),
		UserRpc:    newUserRpc(c.UserRpc),
		ProductRpc: newProductRpc(c.ProductRpc),
	}
}

// 同步表
func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Order{},     // 订单表
		&model.OrderItem{}, // 订单明细表
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}

// 强依赖 product rpc
func newProductRpc(c zrpc.RpcClientConf) pruductpb.ProductClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] product rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] product rpc connection success : %v\n")
	return product.NewProduct(client)
}

// 强依赖 product rpc
func newUserRpc(c zrpc.RpcClientConf) userpb.UserClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] user rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] user rpc connection success : %v\n")
	return user.NewUser(client)
}
