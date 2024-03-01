package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"os"
	"server/app/product/cmd/rpc/internal/config"
	"server/app/product/model"
	"server/common/initialize"
)

type ServiceContext struct {
	Config config.Config

	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := initialize.InitDB(c.Mysql)
	if err != nil {
		logx.Error("init mysql database err", err.Error())
		os.Exit(0)
	}
	rdb, err := initialize.InitRedis(c.CRedis)
	if err != nil {
		logx.Error("init redis database err", err.Error())
		os.Exit(0)
	}

	// autoMigrate(db)

	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(db, rdb),
	}
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Category{},
		&model.Product{},
		&model.FkProduct{},
	)
	if err != nil {
		logx.Error("[DATABASE AutoMigrate ERROR] : ", err)
		os.Exit(0)
	}
	logx.Info("[DATABASE AutoMigrate SUCCESS]")
}
