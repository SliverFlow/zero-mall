package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"os"
	pruductpb "server/app/product/cmd/rpc/pb"
	"server/app/product/cmd/rpc/product"
	"server/app/user/cmd/rpc/internal/config"
	"server/app/user/model"
	"server/common"
)

type ServiceContext struct {
	Config config.Config

	UserModel  model.UserModel
	ProductRpc product.Product
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
		UserModel:  model.NewUserModel(db),
		ProductRpc: newProductRpc(c.ProductRpc),
	}
}

// 同步表
func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},     // 用户表
		&model.Business{}, // 商家表
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
